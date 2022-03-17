package main

import (
	"context"
	"ent/ent"
	"ent/ent/user"
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
	if err = client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources:%v", err)
	}
	//CreateUser(ctx, client, 123, "爱搭搭")
	//CreateUser(ctx, client, 18, "委托维尔")
	q, err := QueryUser(ctx, client)
	if err != nil {
		log.Fatal("query user error: ", err)
	}
	log.Println(q)
}

func CreateUser(ctx context.Context, client *ent.Client,
	age int, name string) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(age).
		SetName(name).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.Query().Where(user.NameEQ("张三")).Only(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("query user: ", u)

	allUsr := client.User.Query().AllX(ctx)
	log.Printf("all user: %+v", allUsr)
	return u, nil
}
