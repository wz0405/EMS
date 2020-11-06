package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbUser     = "admin@admin.com"
	dbPassword = "admin"
	dbName     = "bms"
)

func PostgreSQLHandler(c *gin.Context) {
	// 몽고DB 연결

	clientOptions := options.Client().ApplyURI("mongodb://root:admin1234@db.pms-v2.svc:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

}

// func PostgreSQLHandler(c *gin.Context) {
// 	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
// 		DB_USER, DB_PASSWORD, DB_NAME)
// 	db, err := sql.Open("postgres", dbinfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	result, err := db.Exec("select * from public.bms_alarm_code")
// 	if err != nil {
// 		panic(err)
// 	}
// 	cntAffected, err := result.RowsAffected()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Affected Rows:", cntAffected)
// }
