package storage

import (
	"context"
	"github.com/ioctl/mongoz/types"
	"testing"
)

var ctx = context.Background()

func TestCreate(t *testing.T) {
	InitMongo()
	id, err := ArticleStorage.Create(ctx, &types.Article{
		UserID: "100000000",
		Title:  "mongo document",
		Content: `
Insert a Document
Overview
In this guide, you can learn how to insert documents into a MongoDB collection.

Before you can find, update, and delete documents in MongoDB, you need to insert those documents. You can insert one document using the InsertOne() method, or insert multiple documents using either the InsertMany() or BulkWrite() method.

The following sections focus on InsertOne() and InsertMany().`,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("id: %v\n", id)
}
