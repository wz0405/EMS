package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gopkg.in/mgo.v2/bson"
)

type Graph struct {
	ID  primitive.ObjectID
	num string `json:"id"`
	x   int    `json:"x"`
	y   int    `json:"y"`
}

func GraphHandler(c *gin.Context) {

	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "graph.html", gin.H{
		"title": "메인화면",
	})
}

func DrawGraph(c *gin.Context) {
	c.Header("Content-Type", "application/json charset=utf-8")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions1 := options.Client().ApplyURI("mongodb://localhost:27017")
	client1, err := mongo.NewClient(clientOptions1)
	if err != nil {
		log.Fatal(err)
	}
	err = client1.Connect(context.Background())
	collection1 := client1.Database("local").Collection("graph1")
	//Find
	findOptions := options.Find()
	findOptions.SetProjection(bson.M{
		"_id": 0,
		"id":  1,
		"x":   1,
		"y":   1,
	})
	filter := bson.M{}
	cursor, err := collection1.Find(ctx, filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	//find한 결과를 전부 json으로 변환
	var episodes []bson.M
	cursor.All(ctx, &episodes)

	//find한 결과를 하나씩 꺼내서 사용
	for cursor.Next(ctx) {
		var episode bson.M
		if err = cursor.Decode(&episode); err != nil {
			log.Fatal(err)
		}
		s, _ := json.Marshal(episode)
		fmt.Println(string(s))
	}

	//FindOne
	//find 1개를 한 값을 넘겨줌
	doc1 := collection1.FindOne(context.TODO(), bson.M{})
	var finds bson.M
	doc1.Decode(&finds)

	c.JSON(200, episodes)
}
func DrawGraph1(c *gin.Context) {
	c.Header("Content-Type", "application/json charset=utf-8")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions1 := options.Client().ApplyURI("mongodb://localhost:27017")
	client1, err := mongo.NewClient(clientOptions1)
	if err != nil {
		log.Fatal(err)
	}
	err = client1.Connect(context.Background())
	collection1 := client1.Database("local").Collection("graph2")
	//Find
	findOptions := options.Find()
	findOptions.SetProjection(bson.M{
		"_id": 0,
		"id":  1,
		"x":   1,
		"y":   1,
	})
	filter := bson.M{}
	cursor, err := collection1.Find(ctx, filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	//find한 결과를 전부 json으로 변환
	var episodes []bson.M
	cursor.All(ctx, &episodes)

	//find한 결과를 하나씩 꺼내서 사용
	for cursor.Next(ctx) {
		var episode bson.M
		if err = cursor.Decode(&episode); err != nil {
			log.Fatal(err)
		}
		s, _ := json.Marshal(episode)
		fmt.Println(string(s))
	}

	//FindOne
	//find 1개를 한 값을 넘겨줌
	doc1 := collection1.FindOne(context.TODO(), bson.M{})
	var finds bson.M
	doc1.Decode(&finds)

	c.JSON(200, episodes)
}
func DrawGraph2(c *gin.Context) {
	c.Header("Content-Type", "application/json charset=utf-8")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions1 := options.Client().ApplyURI("mongodb://localhost:27017")
	client1, err := mongo.NewClient(clientOptions1)
	if err != nil {
		log.Fatal(err)
	}
	err = client1.Connect(context.Background())
	collection1 := client1.Database("local").Collection("graph3")

	//Find
	findOptions1 := options.FindOne()

	findOptions1.SetProjection(bson.M{
		"_id": 0,
		"w":   1,
		"z":   1,
		"x":   1,
		"y":   1,
	})
	filter := bson.M{}

	//FindOne
	//find 1개를 한 값을 넘겨줌
	doc1 := collection1.FindOne(ctx, filter, findOptions1)
	var finds bson.M
	doc1.Decode(&finds)
	c.JSON(200, finds)
}
