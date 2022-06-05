package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMongoStoreComments(t *testing.T) {

	client, err := mongo.Connect(context.TODO(),
		options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		t.Fatal("MongoDB 客户端创建失败")
	}
	t.Log("MongoDB client 创建成功")
	// 测试结束后，自动关闭链接
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
		t.Log("MongoDB 关闭成功")
	}(client, context.TODO())
	// 创建一个 database 和一个 collection

	coll := client.Database("test").Collection("test")

	insertOneResult, err := coll.InsertOne(context.TODO(), &Comment{
		Id:      0,
		VideoId: 0,
		CommentUser: User{
			Name: "test",
		},
		Content:    "评论测试",
		CreateDate: "日期",
	})
	if err != nil {
		t.Fatal("Mongodb 插入失败")
	}
	t.Logf("Mongodb 插入成功：%v", insertOneResult.InsertedID)

	comlist := []*Comment{{Id: 0, Content: "acc"}, {Id: 2}}
	// 插入评论队列
	one, err := coll.InsertOne(context.TODO(), &CommentList{
		CommentList: comlist,
	})
	if err != nil {
		t.Fatal("插入 comment list 失败")
	}
	t.Logf("插入 comment list  成功：%v", one.InsertedID)

	// 查找某个特定的 collection
	findResult, err := coll.Find(context.TODO(), bson.D{{"id", 0}})
	if err != nil {
		t.Fatal("未查找到符合的内容")
	}
	t.Logf("查找成功：%v", *findResult)
	var results []*Comment
	_ = findResult.All(context.TODO(), &results)
	if len(results) == 0 {
		t.Fatal("查找失败")
	}
	t.Logf("找到某对象：%v", results[0].Content)
}
