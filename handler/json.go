package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Date struct {
	Id string `json:"Id"`
}

func JsonHandler(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "json.html", gin.H{
		"title": "json화면",
	})
}

func JsonTestHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json charset=utf-8")

	var reqjson Date

	c.Bind(&reqjson)

	// url := "http://106.254.240.205:3030/posts/"
	// numb := reqjson.date
	//url = url + numb
	fmt.Println(reqjson.Id)

	//Get통신
	//data := GetCall(url)

	//json.Unmarshal([]byte(data), &resjson)

	c.JSON(200, reqjson)

}
