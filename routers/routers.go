package routers

import (
	"blog-rest/controller"
	"blog-rest/dbManage"
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() {

	r := gin.Default()

	fmt.Println("=start===")

	articleRoutes := r.Group("/article")
	{
		// add article
		articleRoutes.POST("/add",controller.AddArticles)
	}

	dbManage.InitDb()
	
	r.Run(":9090")
}

