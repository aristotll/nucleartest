package main

import (
	"context"
	"ent/ent"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func CreateOrRefreshTable() {
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
}

func CreateGraph(ctx context.Context, client *ent.Client) error {
	// First, create the users.
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("Ariel").
		Save(ctx)
	if err != nil {
		return err
	}
	neta, err := client.User.
		Create().
		SetAge(28).
		SetName("Neta").
		Save(ctx)
	if err != nil {
		return err
	}

	// Then, create the cars, and attach them to the users in the creation.
	_, err = client.Car.
		Create().
		SetModel("Tesla").
		SetRegisterAt(time.Now()). // ignore the time in the graph.
		SetOwner(a8m).             // attach this graph to Ariel.
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Car.
		Create().
		SetModel("Mazda").
		SetRegisterAt(time.Now()). // ignore the time in the graph.
		SetOwner(a8m).             // attach this graph to Ariel.
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Car.
		Create().
		SetModel("Ford").
		SetRegisterAt(time.Now()). // ignore the time in the graph.
		SetOwner(neta).            // attach this graph to Neta.
		Save(ctx)
	if err != nil {
		return err
	}
	// Create the groups, and add their users in the creation.
	_, err = client.Group.
		Create().
		SetName("GitLab").
		AddUsers(neta, a8m).
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Group.
		Create().
		SetName("GitHub").
		AddUsers(a8m).
		Save(ctx)
	if err != nil {
		return err
	}
	log.Println("The graph was created successfully")
	return nil
}

func main() {
	client, err := ent.Open("mysql",
		"root:rootroot@tcp(127.0.0.1:3306)/keepalive?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	ctx := context.Background()
	//CreateOrRefreshTable()
	CreateGraph(ctx, client)
}
