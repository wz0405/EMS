package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

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
	client := http.Client{Timeout: 5 * time.Second}
	c.Header("Content-Type", "application/json charset=utf-8")

	var reqjson Date

	c.Bind(&reqjson)

	url := "http://106.254.240.205:3030/posts/"
	numb := reqjson.Id
	url = url + numb
	//d := c.PostForm(url)
	var data Board
	//Get통신
	resp, _ := client.Get(url)
	respBody, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(respBody), &data)

	c.JSON(200, data)

}
