package main

import (
	handler "web.com/handler"
)

func (server *serverStruct) SetRoute() bool {

	//TODO: 아래 부분에 route할 uri와 핸들러들을 추가한다.
	server.route.LoadHTMLGlob("templates/*")

	server.route.GET("/login", handler.LoginHandler)
	server.route.POST("/login_action", handler.LoginActionHandler)

	server.route.GET("/main", handler.MainHandler)
	server.route.POST("/main_viewData", handler.MainViewDataHandler)

	return true
}
