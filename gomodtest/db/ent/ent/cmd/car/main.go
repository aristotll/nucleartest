package main

import (
	"context"
	"ent/ent"
	"ent/ent/user"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	client, err := ent.Open("mysql",
		"root:rootroot@tcp(127.0.0.1:3306)/keepalive?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	ctx := context.Background()
	//CreateCar(ctx, client)
	QueryCars(ctx, client)
	QueryCarUsers(ctx, client)
}

func CreateCar(ctx context.Context, client *ent.Client) (*ent.User, error) {
	tesla, err := client.Car.
		Create().
		SetModel("Tesla").
		SetRegisterAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", tesla)

	ford, err := client.Car.
		Create().
		SetModel("Ford").
		SetRegisterAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", ford)

	a8m, err := client.User.
		Create().
		SetName("a8m").
		SetAge(30).
		AddCars(tesla, ford).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", a8m)
	return a8m, nil
}

// 一对多查询：一个用户有多辆汽车
func QueryCars(ctx context.Context, client *ent.Client) error {
	a8m, err := client.User.
		Query().
		Where(user.NameEQ("a8m")).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("query user [a8m] error: %w", err)
	}

	cars := a8m.QueryCars().AllX(ctx)
	log.Println("cars: ", cars)
	return nil
}

// 反向查询，查出车的所有者
func QueryCarUsers(ctx context.Context, client *ent.Client) error {
	a8m, err := client.Debug().User.
		Query().
		Where(user.NameEQ("a8m")).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("query user [a8m] error: %w", err)
	}

	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("query all cars error: %w", err)
	}

	for _, car := range cars {
		owner, err := car.QueryOwner().Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying car %q owner: %w", car.Model, err)
		}
		log.Printf("car %q owner: %q\n", car.Model, owner.Name)
	}
	return nil
}
