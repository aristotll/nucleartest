package main

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"io/ioutil"
)

type Struct struct {
	Mysql struct {
		Username string
		Password string
	}

	CloudAdmin struct {
		Username string
		Password string
	}

	CloudAddr string
	NasAddr   string
}

func main() {
	f, err := os.Open("./config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	//b := make([]byte, 1024)
	//_, err = f.Read(b)
	//if err != nil {
	//	log.Fatalln("read file error: ", err)
	//}

	b, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatalln("read file error: ", err)
	}

	s := &Struct{}
	if err := yaml.Unmarshal(b, s); err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v\n", s)
}
