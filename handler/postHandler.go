package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiPostHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json charset=utf-8")
	book := Board{100, "id", "pwd"}
	pbytes, _ := json.Marshal(book)
	buff := bytes.NewBuffer(pbytes)
	response, err := http.Post("http://106.254.240.205:3030/posts/", "application/json", buff)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	// Response 체크.
	rBody, err := ioutil.ReadAll(response.Body)
	if err == nil {
		str := string(rBody)
		println(str)
	}
}
func AddProductHandler(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "toy.html", gin.H{
		"title": "toy",
	})
}
func AddProduct(c *gin.Context) {
	frmCategory := c.PostForm("key")
	fmt.Println(frmCategory)
}
func PostTest(c *gin.Context) {
	c.Header("Content-Type", "application/json charset=utf-8")
	var jsons Board
	err := json.NewDecoder(c.Request.Body).Decode(&jsons)
	if err != nil { 
		fmt.Println(err)
	}
	jsons.Title = "hey"
	jsons.Content = "hi"

	fmt.Println(jsons)
	c.JSON(200, jsons)
}
