package parsexml

import (
	"testing"
)

func TestParseCatalogFile(t *testing.T) {
	zr := NewZipReader("./成绩单_正确档案包.zip")
	f, err := zr.GetFileBytes("目录文件.xml")
	if err != nil {
		panic(err)
	}

	b := NewXmlCheck()
	if err != nil {
		panic(err)
	}
	if err := b.checkCatalogFile(f); err != nil {
		panic(err)
	}
}

//func TestParseArchiveInformation(t *testing.T) {
//	zr := NewZipReader("./成绩单_正确档案包.zip")
//	f, err := zr.GetFileBytes("电子档案归档信息表.xml")
//	if err != nil {
//		panic(err)
//	}
//
//	b, err := NewXmlParseFromBytes(f)
//	if err != nil {
//		panic(err)
//	}
//	b.ParseArchiveInformation()
//}
//
//func TestParseMetaData(t *testing.T) {
//	zr := NewZipReader("./成绩单_正确档案包.zip")
//	f, err := zr.GetFileBytes("成绩单.xml")
//	if err != nil {
//		panic(err)
//	}
//
//	b, err := NewXmlParseFromBytes(f)
//	if err != nil {
//		panic(err)
//	}
//	b.ParseMetaData()
//}
