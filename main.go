package main

import (
	"fmt"
	"love-comme-server/routers"
	"love-comme-server/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	api := gin.Default()
	utils.LoadEnv()
	routers.InitRouter(api)
	api.Run(fmt.Sprintf(":%s", utils.ApiPort))
}
