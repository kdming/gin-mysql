package main

import (
	"app/api"
	"app/common/config"
	"app/dao/mysql"
	"github.com/gin-gonic/gin"
)

func main() {
	mysql.Connect()
	g := gin.New()
	api.InitRouter(g)
	appPort := config.GetConfig().AppPort
	if appPort == "" {
		appPort = ":8080"
	}
	if err := g.Run(config.GetConfig().AppPort); err != nil {
		panic(err)
	}
}
