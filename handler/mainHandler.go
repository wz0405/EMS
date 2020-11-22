package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	value struct {
		Search string `json:"search"`
	}
)

type ()

func MainHandler(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "main.html", gin.H{
		"title": "메인화면",
	})
}

func MainViewDataHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json charset=utf-8")

	aValue := "Apple"
	bValue := "12345"
	cValue := "D는 현재시간값입니다"
	dVaule := time.Now().Format("2006-01-02 15:04:05")

	c.JSON(200, gin.H{
		"A": aValue,
		"B": bValue,
		"C": cValue,
		"D": dVaule,
	})

}

func ReceiveData(c *gin.Context) {
	//json형태로 값을 받음
	data, _ := c.GetRawData()

	reqBody := value{}
	err := json.Unmarshal([]byte(data), &reqBody)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	resp, _ := http.Get("http://106.254.240.205:3030/posts/" + reqBody.Search)

	respBody, _ := ioutil.ReadAll(resp.Body)

	var data1 Board
	json.Unmarshal([]byte(respBody), &data1)

	c.JSON(200, data1)

}
