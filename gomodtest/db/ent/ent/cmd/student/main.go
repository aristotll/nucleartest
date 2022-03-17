package main

import (
	"context"
	"ent/ent"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	client, err := ent.Open("mysql",
		"root:rootroot@tcp(127.0.0.1:3306)/keepalive?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	ctx := context.Background()
	QueryStudent(ctx, client)
}

func QueryStudent(ctx context.Context, client *ent.Client) {
	students, err := client.Debug().Student.Query().All(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(students)
}
