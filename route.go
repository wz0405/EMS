package main

import (
	handler "web.com/handler"
)

func (server *serverStruct) SetRoute() bool {

	//TODO: 아래 부분에 route할 uri와 핸들러들을 추가한다.
	server.route.LoadHTMLGlob("templates/*")
	server.route.Static("/assets", "./assets")

	server.route.GET("/login", handler.LoginHandler)
	server.route.GET("/login2", handler.LoginHandler2)
	server.route.POST("/login_action", handler.LoginActionHandler)
	server.route.POST("/login_action2", handler.LoginActionHandler2)

	server.route.GET("/main", handler.MainHandler)
	server.route.POST("/main_viewData", handler.MainViewDataHandler)
	server.route.POST("/api", handler.APIBlogSearchHandler)
	server.route.POST("/sendToGo", handler.ReceiveData)

	server.route.GET("/db", handler.PostgreSQLHandler)

	server.route.GET("/post", handler.ApiPostHandler)

	server.route.GET("/add", handler.AddProductHandler)
	server.route.POST("/addProduct", handler.AddProduct)

	server.route.POST("/postTest", handler.PostTest)

	server.route.GET("/graph", handler.GraphHandler)
	server.route.GET("/getGraphData", handler.DrawGraph)
	server.route.GET("/getGraphData1", handler.DrawGraph1)
	server.route.GET("/getGraphData2", handler.DrawGraph2)

	server.route.GET("/jsonTest", handler.JsonHandler)
	server.route.POST("/json_test", handler.JsonTestHandler)
	return true
}
