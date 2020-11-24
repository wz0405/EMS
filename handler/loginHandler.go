package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	loginRequestData struct {
		LoginID string `json:"loginID"`
		LoginPW string `json:"loginPW"`
	}
)

func LoginHandler(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "로그인화면",
	})
}
func LoginHandler2(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "web0001.html", gin.H{
		"title": "로그인화면",
	})
}

func LoginActionHandler(c *gin.Context) {
	//TODO: 하드코딩으로 고정된 아이디와 비밀번호임. 추후 DB에 맞게 변경
	c.Header("Content-Type", "application/json charset=utf-8")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions1 := options.Client().ApplyURI("mongodb://localhost:27017")
	client1, err := mongo.NewClient(clientOptions1)
	if err != nil {
		log.Fatal(err)
	}
	err = client1.Connect(context.Background())
	collection1 := client1.Database("local").Collection("login")
	doc1 := collection1.FindOne(ctx, bson.M{})
	var login loginRequestData
	//var settings Settings
	doc1.Decode(&login)
	reqBody := loginRequestData{}

	data, err := c.GetRawData()

	//	dataString := string(data)
	//	fmt.Println(dataString)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = json.Unmarshal([]byte(data), &reqBody)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	id := login.LoginID
	pw := login.LoginPW

	fmt.Println("Received ID:" + reqBody.LoginID + " PW:" + reqBody.LoginPW)

	loginSuccess := false
	if id == reqBody.LoginID {
		if pw == reqBody.LoginPW {
			//로그인 처리하기
			loginSuccess = true
		}
	}

	if loginSuccess {
		session := sessions.Default(c)
		session.Set(reqBody.LoginID, reqBody.LoginPW) //TODO: 로그인된 유저의 세션정보 확인하기

		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
			return
		}

	}

	c.JSON(200, gin.H{
		"result": loginSuccess,
	})
}
func LoginActionHandler2(c *gin.Context) {
	//TODO: 하드코딩으로 고정된 아이디와 비밀번호임. 추후 DB에 맞게 변경
	c.Header("Content-Type", "application/json charset=utf-8")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions1 := options.Client().ApplyURI("mongodb://localhost:27017")
	client1, err := mongo.NewClient(clientOptions1)
	if err != nil {
		log.Fatal(err)
	}
	err = client1.Connect(context.Background())
	collection1 := client1.Database("local").Collection("login")
	doc1 := collection1.FindOne(ctx, bson.M{})
	var login loginRequestData
	//var settings Settings
	doc1.Decode(&login)
	reqBody := loginRequestData{}

	data, err := c.GetRawData()

	//	dataString := string(data)
	//	fmt.Println(dataString)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = json.Unmarshal([]byte(data), &reqBody)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	id := login.LoginID
	pw := login.LoginPW

	fmt.Println("Received ID:" + reqBody.LoginID + " PW:" + reqBody.LoginPW)

	loginSuccess := false
	if id == reqBody.LoginID {
		if pw == reqBody.LoginPW {
			//로그인 처리하기
			loginSuccess = true
		}
	}

	if loginSuccess {
		session := sessions.Default(c)
		session.Set(reqBody.LoginID, reqBody.LoginPW) //TODO: 로그인된 유저의 세션정보 확인하기

		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
			return
		}

	}

	c.JSON(200, gin.H{
		"result": loginSuccess,
	})
}
