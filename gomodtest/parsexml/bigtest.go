package parsexml

import (
	"archive/zip"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/beevik/etree"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

type XmlCheck struct {
	doc       *etree.Document
	fileCount int64
}

func NewXmlCheck() *XmlCheck {
	return &XmlCheck{doc: etree.NewDocument()}
}

func NewXmlCheckFromFile(src string) (*XmlCheck, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(src); err != nil {
		return nil, err
	}

	return &XmlCheck{doc: doc}, nil
}

func NewXmlCheckFromBytes(b []byte) (*XmlCheck, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(b); err != nil {
		return nil, err
	}

	return &XmlCheck{doc: doc}, nil
}

// CheckCatalogFile 校验目录文件.xml
func (a *XmlCheck) checkCatalogFile(b []byte) error {
	if err := a.doc.ReadFromBytes(b); err != nil {
		return err
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
func (a *XmlCheck) checkMetaData(src string) error {
	if err := a.doc.ReadFromFile(src); err != nil {
		return err
	}
	ele := a.doc.FindElement("./文件元数据信息/类型相关元数据/原文件")
	for _, e := range ele.ChildElements() {
		fmt.Println(e.Tag, e.Text())
	}
	fmt.Println("===========================")

	ele = a.doc.FindElement("./文件元数据信息/类型相关元数据/轻量级OFD")
	for _, e := range ele.ChildElements() {
		fmt.Println(e.Tag, e.Text())
	}
	fmt.Println("===========================")

	ele = a.doc.FindElement("./文件元数据信息/类型相关元数据/重量级OFD")
	for _, e := range ele.ChildElements() {
		fmt.Println(e.Tag, e.Text())
	}

	return nil
}

// CheckArchiveInformation 解析电子档案归档信息表.xml
func (a *XmlCheck) checkArchiveInformation(b []byte) error {
	if err := a.doc.ReadFromBytes(b); err != nil {
		return err
	}
	root := a.doc.SelectElement("xml")
	for _, e := range root.ChildElements() {
		if e.Tag == "FILECOUNT" {
			vs, _ := strconv.Atoi(e.Text())
			if int64(vs) != a.fileCount {
				return errors.New("数量不符")
			}
		}
		fmt.Println(e.Tag, e.Text())
	}
	return nil
}

func (a *XmlCheck) Check(b []byte) error {
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
		if strings.Contains(file.Name, filename) {
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
		if strings.Contains(file.Name, filename) {
			return file, nil
		}
	}
	return nil, nil
}

func GetMd5(src string) (string, error) {
	h := md5.New()
	f, err := os.Open(src)
	if err != nil {
		return "", err
	}
	io.Copy(h, f)
	v := hex.EncodeToString(h.Sum(nil))
	return v, nil
}

// CheckMetaData 检查元数据对应的 xml
func (z *ZipReader) CheckMetaData(src string, md5, name, format string, size int64) error {
	f, err := z.GetFile(src)
	if err != nil {
		return err
	}

	if f.FileInfo().Name() != name {
		return errors.New(fmt.Sprintf("check file %s error: 电子文件名与元数据不相符", name))
	}

	if f.FileInfo().Size() != size {

	}

	if path.Ext(name) != format {

	}

	md55, err := GetMd5(src)
	if err != nil {
		return err
	}

	if md55 != md5 {

	}

	return nil
}

func (z *ZipReader) CheckCatalogFile() {

}

type Archives struct{}
