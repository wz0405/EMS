package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

func (server *serverStruct) init() {
	config := new(serverConfig)
	confFile, err := ioutil.ReadFile("config.yaml")

	err = yaml.Unmarshal(confFile, &config)
	if err != nil {
		//TODO: 오류 메세지 출력하기
		return
	}

	server.config = config
	server.route = gin.Default()
	store := cookie.NewStore([]byte("secret"))
	server.route.Use(sessions.Sessions("mysession", store))

	server.SetRoute()

	server.server = &http.Server{
		Addr:           ":8080",
		Handler:        server.route,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

}

func (server *serverStruct) startServer() {
	server.server.ListenAndServe()
}
