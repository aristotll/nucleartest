package parsexml

import (
	"fmt"
	"github.com/beevik/etree"
)

func parse() {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("./x.xml"); err != nil {
		panic(err)
	}

	root := doc.SelectElement("bookstore")
	for _, book := range root.SelectElements("book") {
		//fmt.Println("child element: ", book.Tag)
		if title := book.SelectElement("title"); title != nil {
			// 未找到则返回 unknown
			lang := title.SelectAttrValue("lang", "unknown")
			fmt.Printf(" TITLE: %s (%s)\n", title.Text(), lang)
		}

		for _, attr := range book.Attr {
			fmt.Printf(" ATTR: %s = %s\n", attr.Key, attr.Value)
		}
	}

	fmt.Println("search: ")
	for _, t := range doc.FindElements("//book[@category='WEB']/title") {
		fmt.Println(" Title:", t.Text())
	}

}

func ParseArchives() {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("./成绩单.xml"); err != nil {
		panic(err)
	}

	for _, e := range doc.FindElements("./文件元数据信息/类型相关元数据/原文件/*") {
		fmt.Printf("%s: %s\n", e.Tag, e.Text())
	}

	fmt.Println("=============================")

	for _, e := range doc.FindElements("./文件元数据信息/类型相关元数据/轻量级OFD/*") {
		fmt.Printf("%s: %s\n", e.Tag, e.Text())
	}
}

func parse1() {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("./目录文件.xml"); err != nil {
		panic(err)
	}

	ele := doc.FindElements("./目录文件/文件目录/*")
	//fmt.Printf("%+v \n", ele)
	for _, e := range ele {
		for _, ee := range e.ChildElements() {
			fmt.Println(ee.Tag, ee.Text())
			if ee.Tag == "" {

			}
		}
	}

}
