package handler

import (
	"github.com/gin-gonic/gin"
)

type ()

func APIBlogSearchHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json charset=utf-8")

	//TODO: 해당주소에 API호출 프로토콜 가져와보기
	//response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")

}
