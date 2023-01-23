package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoClint(ctx context.Context, url string) (client *mongo.Client, err error) {
	client, err = mongo.Connect(ctx,
		options.Client().ApplyURI(url),
		options.Client().SetConnectTimeout(time.Second*5))
	if err != nil {
		return nil, err
	}
	//defer func() {
	//	if err = client.Disconnect(ctx); err != nil {
	//		return
	//	}
	//}()
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return client, nil
}

func InitMongo() {
	ctx := context.Background()
	mongoURL := "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.6.0"
	mongoDatabase := "test"
	cli, err := NewMongoClint(ctx, mongoURL)
	if err != nil {
		panic(err)
	}
	initStorage(ctx, cli, mongoDatabase)
}

func initStorage(ctx context.Context, cli *mongo.Client, mongoDatabase string) {
	var err error
	ArticleStorage, err = newArticleStorage(ctx, cli, mongoDatabase)
	if err != nil {
		panic(err)
	}
}
