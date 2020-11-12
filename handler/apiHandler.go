package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Board struct {
	ID      int    `json:"ID"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
}

func APIBlogSearchHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json charset=utf-8")

	//TODO: 해당주소에 API호출 프로토콜 가져와보기
	resp, err := http.Get("http://106.254.240.205:3030/posts/")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	var data map[string]interface{}
	var datas []Board
	json.Unmarshal([]byte(respBody), &data)
	json.Unmarshal([]byte(respBody), &datas)
	if err == nil {
		//str := string(respBody)
		//println(str)
		//fmt.Printf("%s\n", str)
	}

	fmt.Println(data["id"], data["title"])
	idValue := datas[0].ID
	titleValue := datas[0].Title
	contentValue := datas[0].Content

	c.JSON(200, gin.H{
		"E": idValue,
		"F": titleValue,
		"G": contentValue,
	})
}
