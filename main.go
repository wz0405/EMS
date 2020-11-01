package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type serverStruct struct {
	route  *gin.Engine
	server *http.Server
	config *serverConfig
}

type serverConfig struct {
	port string `yaml:"Port"`
}

func main() {

	server := new(serverStruct)
	server.init()

	fmt.Println("Start Server!")
	server.startServer()
}
