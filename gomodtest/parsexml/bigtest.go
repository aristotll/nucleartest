package parsexml

import (
	"archive/zip"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/beevik/etree"
	"io"
	"path"
	"strconv"
)

type XmlChecker struct {
	doc        *etree.Document
	zipReader  *ZipReader
	fileCount  int64
	efileCount int64 // <原文件> <轻量级OFD> <重量级OFD> 总数量
	totalSize  int64 // <原文件> <轻量级OFD> <重量级OFD> 总大小
}

func NewXmlChecker(zipPath string) *XmlChecker {
	return &XmlChecker{
		doc:       etree.NewDocument(),
		zipReader: NewZipReader(zipPath),
	}
}

func (a *XmlChecker) Check() error {
	// 该方法必须先调用
	if err := a.checkCatalogFile(); err != nil {
		return err
	}
	if err := a.checkArchiveInformation(); err != nil {
		return err
	}
	return nil
}

// CheckCatalogFile 校验目录文件.xml
func (a *XmlChecker) checkCatalogFile() error {
	b, err := a.zipReader.GetFileBytes("档案包/目录文件.xml")
	if err != nil {
		return fmt.Errorf("打开 [目录文件.xml] 失败：%v", err.Error())
	}

	if err := a.doc.ReadFromBytes(b); err != nil {
		return fmt.Errorf("读取 [目录文件.xml] 失败：%v", err.Error())
	}
	var index = 1
	ele := a.doc.FindElements("./目录文件/文件目录/*")
	a.fileCount = int64(len(ele))
	for _, e := range ele {
		for _, ee := range e.ChildElements() {
			if ee.Tag == "顺序号" {
				ind, _ := strconv.Atoi(ee.Text())
				if ind != index {
					return errors.New("顺序号错误")
				}
				index++
			}
			if ee.Tag == "元数据路径" {
				if err := a.checkMetaData(ee.Text()); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// CheckMetaData 解析元数据对应的 xml
func (a *XmlChecker) checkMetaData(src string) error {
	b, err := a.zipReader.GetFileBytes(src)
	if err != nil {
		return err
	}

	// 坑：这里必须重新 new 一个，否则解析会出问题，虽然指定了 ReadFromBytes(b)，但读取的依然是
	// 之前的 xml，造成解析异常
	a.doc = etree.NewDocument()
	if err := a.doc.ReadFromBytes(b); err != nil {
		return err
	}

	eles := a.doc.FindElement("./文件元数据信息/类型相关元数据")
	for _, ele := range eles.ChildElements() {
		// 如果当前标签的子标签中有 <包内路径>，则表示这是一条文件信息相关标签，比如 <原文件>、<重量级OFD>
		// 不同类型的档案，文件信息标签名可能不同，比如学籍档案中，多出了 <入学照片> 这一标签
		// 论文档案中，没有 <原文件> 这一标签，而是 <图书馆审核后论文>，因为不同类型的档案的标签名不统一，
		// 所以不能通过固定标签名去获取信息
		if v := ele.FindElement("包内路径"); v != nil {
			if err := a.check(ele); err != nil {
				return err
			}
		}
	}

	//ele := a.doc.FindElement("./文件元数据信息/类型相关元数据/原文件")
	//if ele != nil {
	//	if err := a.check(ele); err != nil {
	//		return err
	//	}
	//}
	//
	//ele = a.doc.FindElement("./文件元数据信息/类型相关元数据/轻量级OFD")
	//if ele != nil {
	//	if err := a.check(ele); err != nil {
	//		return err
	//	}
	//}
	//
	//ele = a.doc.FindElement("./文件元数据信息/类型相关元数据/重量级OFD")
	//if ele != nil {
	//	if err := a.check(ele); err != nil {
	//		return err
	//	}
	//}
	//
	//ele = a.doc.FindElement("./文件元数据信息/类型相关元数据/签章文件")
	//if ele != nil {
	//	if err := a.check(ele); err != nil {
	//		return err
	//	}
	//}

	return nil
}

// CheckArchiveInformation 解析 电子档案归档信息表.xml
func (a *XmlChecker) checkArchiveInformation() error {
	b, err := a.zipReader.GetFileBytes("档案包/其他/电子档案归档信息表.xml")
	if err != nil {
		return fmt.Errorf("打开 [电子档案归档信息表.xml] 失败：%v", err.Error())
	}

	if err := a.doc.ReadFromBytes(b); err != nil {
		return err
	}
	root := a.doc.SelectElement("xml")
	for _, e := range root.ChildElements() {
		if e.Tag == "FILECOUNT" {
			vs, _ := strconv.Atoi(e.Text())
			if int64(vs) != a.fileCount {
				return fmt.Errorf("数量(FILECOUNT)不符，当前：%v，xml 中：%v", a.fileCount, vs)
			}
		}
		if e.Tag == "EFILECOUNT" {
			vs, _ := strconv.Atoi(e.Text())
			if int64(vs) != a.efileCount {
				return fmt.Errorf("数量(EFILECOUNT)不符，当前：%v，xml 中：%v", a.efileCount, vs)
			}
		}
		if e.Tag == "TOTALSIZE" {
			vs, _ := strconv.Atoi(e.Text())
			if int64(vs) != a.totalSize {
				return fmt.Errorf("总大小(TOTALSIZE)不符，当前：%v，xml 中：%v", a.totalSize, vs)
			}
		}
	}
	return nil
}

func (a *XmlChecker) GetMd5(src string) (string, error) {
	b, err := a.zipReader.GetFileBytes(src)
	if err != nil {
		return "", err
	}

	h := md5.New()
	io.Copy(h, bytes.NewBuffer(b))
	v := hex.EncodeToString(h.Sum(nil))
	return v, nil
}

// CheckMetaData 检查元数据对应的 xml
func (a *XmlChecker) check(ele *etree.Element) error {
	pkgPath := ele.FindElement("包内路径").Text()
	format := ele.FindElement("文件格式").Text()
	filename := ele.FindElement("文件名称").Text()
	filesize := ele.FindElement("文件大小").Text()
	filemd5 := ele.FindElement("MD5").Text()
	filesize_, _ := strconv.Atoi(filesize)

	f, err := a.zipReader.GetFile(pkgPath)
	if err != nil {
		return fmt.Errorf("打开 %v 失败：%v", pkgPath, err)
	}
	if f == nil {
		return fmt.Errorf("打开 %v 失败：文件不存在", pkgPath)
	}

	if f.FileInfo().Name() != filename {
		return fmt.Errorf("check file %s error: 电子文件名与元数据不相符", filename)
	}

	if f.FileInfo().Size() != int64(filesize_) {
		return fmt.Errorf("[%v] 文件大小不匹配", pkgPath)
	}

	ext := path.Ext(f.FileInfo().Name())[1:] // 返回的格式为 .pdf，需要去掉 .
	if ext != format {
		return fmt.Errorf("[%v] 后缀名不匹配 (%v != %v)", pkgPath, ext, format)
	}

	trueMd5, err := a.GetMd5(pkgPath)
	if err != nil {
		return fmt.Errorf("[%v] 获取文件 md5 错误：%v", pkgPath, err)
	}

	if trueMd5 != filemd5 {
		return fmt.Errorf("[%v] md5 不匹配", pkgPath)
	}

	a.efileCount++
	a.totalSize += f.FileInfo().Size()

	return nil
}

type ZipReader struct {
	src string
}

func NewZipReader(src string) *ZipReader {
	return &ZipReader{src: src}
}

// GetFileBytes 从压缩包中获取名为 filename 的文件，以字节形式返回
func (z *ZipReader) GetFileBytes(filename string) ([]byte, error) {
	r, err := zip.OpenReader(z.src)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	for _, file := range r.File {
		// error: strings.Contains(file.Name, filename)
		if file.Name == filename {
			f, err := file.Open()
			if err != nil {
				return nil, err
			}

			b, err := io.ReadAll(f)
			if err != nil {
				return nil, err
			}
			return b, nil
		}
	}
	return nil, nil
}

// GetFile 从压缩包中获取名为 filename 的文件，以文件形式返回
func (z *ZipReader) GetFile(filename string) (*zip.File, error) {
	r, err := zip.OpenReader(z.src)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	for _, file := range r.File {
		if file.Name == filename {
			return file, nil
		}
	}
	return nil, nil
}
