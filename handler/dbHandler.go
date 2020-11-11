package handler

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbUser     = "admin@admin.com"
	dbPassword = "admin1234"
	dbName     = "bms"
	dbIp       = "db.pms-v2.svc:27017"
)

type Dog struct {
	ID     primitive.ObjectID
	Name   string
	Family string
	Age    int
	Weight int
}
type Settings struct {
	ID                         primitive.ObjectID
	Scale                      string
	DefaultResultView          string
	MongoBinaryPath            string
	MaxAllowedFetchSize        int
	AutoCompleteSamplesCount   int
	SocketTimeoutInSeconds     int
	ConnectionTimeoutInSeconds int
	DbStatsScheduler           int
	ShowDBStats                bool
	ShowLiveChat               bool
	Updates                    bool
	SingleTabResultSets        bool
	MaxLiveChartDataPoints     int
	IsMigrationDone            bool
}

func PostgreSQLHandler(c *gin.Context) {
	// db, err := sql.Open("mysql", "root:charger!#123@(localhost:3306)/ship")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// var user_id string
	// err = db.QueryRow("SELECT user_id FROM user where seq=1").Scan(&user_id)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(user_id)

	//몽고DB 연결
	clientOptions1 := options.Client().ApplyURI("mongodb://localhost:27017/local")
	clientOptions := options.Client().ApplyURI("mongodb://db.pms-v2.svc:27017")
	client, err := mongo.NewClient(clientOptions)
	client1, _ := mongo.NewClient(clientOptions1)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	err = client1.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	collection1 := client1.Database("local").Collection("test")
	collection := client.Database("strapi").Collection("settings")
	// find documents

	doc := collection.FindOne(context.TODO(), bson.M{})
	doc1 := collection1.FindOne(context.TODO(), bson.M{})
	var dog Dog
	var settings Settings
	doc1.Decode(&dog)
	fmt.Println(dog)
	doc.Decode(&settings)
	fmt.Println(settings)

}
