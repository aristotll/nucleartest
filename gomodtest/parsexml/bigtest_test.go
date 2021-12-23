package parsexml

import (
	"fmt"
	"github.com/beevik/etree"
	"log"
	"testing"
)

func init() {
	//log.SetFlags(log.Lshortfile)
}

func TestZipReaderGetFile(t *testing.T) {
	zr := NewZipReader("./成绩单_正确档案包.zip")
	f, err := zr.GetFile("档案包/电子档案/6/02教学/CJD/1996/土木系/960113/成绩单.xml")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(f)
}

func Test1(t *testing.T) {
	zr := NewZipReader("./成绩单_正确档案包.zip")
	f, err := zr.GetFileBytes("档案包/电子档案/6/02教学/CJD/1996/土木系/960113/成绩单.xml")
	if err != nil {
		log.Fatalln(err)
	}
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(f); err != nil {
		log.Fatalln(err)
	}
	e := doc.FindElement("./文件元数据信息/类型相关元数据/轻量级OFD")
	ee := e.FindElement("包内路径")
	fmt.Println(ee.Tag, ee.Text())
}

func Test11(t *testing.T) {
	zr := NewZipReader("./论文_正确档案包.zip")
	f, err := zr.GetFileBytes("论文.xml")
	if err != nil {
		log.Fatalln(err)
	}
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(f); err != nil {
		log.Fatalln(err)
	}

	eles := doc.FindElement("./文件元数据信息/类型相关元数据")
	for _, e := range eles.ChildElements() {
		if v := e.FindElement("包内路径"); v != nil {
			fmt.Println("ok")
		}
	}
}

func Test(t *testing.T) {
	paths := []string{
		"./学籍卡_文件MD5不一致.zip",
		"./学籍卡_文件大小不一致.zip",
		"./学籍卡_正确档案包.zip",
		"./学籍卡_文件路径不一致.zip",
		"./学籍卡_文件数量不一致.zip",
		"./论文_文件MD5不一致.zip",
		"./论文_文件大小不一致.zip",
		"./论文_文件数量不一致.zip",
		"./论文_文件路径不一致.zip",
		"./论文_正确档案包.zip",
		"./成绩单_文件MD5不一致.zip",
		"./成绩单_文件大小不一致.zip",
		"./成绩单_文件数量不一致.zip",
		"./成绩单_文件路径不一致.zip",
		"./成绩单_正确档案包.zip",
	}

	for _, p := range paths {
		x := NewXmlChecker(p)
		if err := x.Check(); err != nil {
			log.Printf("验证 [%v] 错误: %v \n", p, err)
		} else {
			log.Printf("验证 [%v] 正确 \n", p)
		}
	}
}

func TestSpecialCase(t *testing.T) {
	x := NewXmlChecker("./学籍卡_正确档案包.zip")
	if err := x.Check(); err != nil {
		log.Fatalln(err)
	}
}

func Test111(t *testing.T) {
	b, err := NewZipReader("./学籍卡_正确档案包.zip").GetFileBytes("档案包/电子档案/6/02教学/XJK/1996/计算机系/961304/学籍卡.xml")
	if err != nil {
		log.Fatalln(err)
	}

	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(b); err != nil {
		log.Fatalln(err)
	}

	ele := doc.FindElement("./文件元数据信息/类型相关元数据")
	for _, e := range ele.ChildElements() {
		if v := e.FindElement("包内路径"); v != nil {
			log.Printf("ok \n")
		}
	}
}
