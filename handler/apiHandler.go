package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Board struct {
	ID      string
	Title   string
	Content string
}

func APIBlogSearchHandler(c *gin.Context) {
	//c.Header("Content-Type", "application/json charset=utf-8")

	//TODO: 해당주소에 API호출 프로토콜 가져와보기
	resp, err := http.Get("http://106.254.240.205:3030/posts/1")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	var data map[string]interface{}

	json.Unmarshal([]byte(respBody), &data)
	// fmt.Println(bo.ID, bo.Title, bo.Content)

	if err == nil {
		str := string(respBody)
		//println(str)
		fmt.Printf("%s\n", str)
	}

	fmt.Println(data["id"], data["title"])
}
