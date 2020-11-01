package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
