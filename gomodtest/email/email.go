package main

import (
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func main() {
	code := "nfjvvphbbadd"
	username := "@qq.com"
	
	e := email.NewEmail()
	e.From = username
	e.To = []string{"zh1105336755@gmail.com"}
	e.Subject = "go test"
	e.Text = []byte("go send email test")
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", username, code, "smtp.qq.com"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("send success")
}
