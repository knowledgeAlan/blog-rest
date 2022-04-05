package main

import (
	"blog-rest/routers"

	"github.com/gin-gonic/gin"
)

func main(){
	gin.SetMode(gin.DebugMode)

	routers.InitializeRoutes()
	
}