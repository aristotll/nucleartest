package storage

import (
	"context"
	"fmt"

	"github.com/ioctl/mongoz/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ArticleStorage *articleStorage

type articleStorage struct {
	Cli        *mongo.Client
	Collection *mongo.Collection // Collection 集合是一组文档，类似于关系型数据库的表
	database   string
}

func newArticleStorage(ctx context.Context, cli *mongo.Client, database string) (*articleStorage, error) {
	s := &articleStorage{
		Cli:        cli,
		Collection: cli.Database(database).Collection(ArticleCollection),
		database:   database,
	}
	if err := s.EnsureIndex(ctx); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *articleStorage) EnsureIndex(ctx context.Context) error {
	index := mongo.IndexModel{
		Keys: bson.M{
			"UserID": 1,
		},
		Options: &options.IndexOptions{},
	}
	_, err := s.Collection.Indexes().CreateOne(ctx, index)
	return err
}

func (s *articleStorage) Create(ctx context.Context, article *types.Article) (string, error) {
	resp, err := s.Collection.InsertOne(ctx, article)
	if err != nil {
		return "", err
	}
	return resp.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (s *articleStorage) Get(ctx context.Context, id string) (*types.Article, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	resp := s.Collection.FindOne(ctx, filter)
	if err := resp.Err(); err != nil {
		return nil, err
	}
	var obj = &types.Article{}
	if err := resp.Decode(obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *articleStorage) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	resp, err := s.Collection.DeleteOne(ctx, objID)
	if err != nil {
		return err
	}
	if resp.DeletedCount == 0 {
		return fmt.Errorf("%v not found", id)
	}
	return nil
}

func (s *articleStorage) Update(ctx context.Context, new *types.Article) error {
	return nil
}
