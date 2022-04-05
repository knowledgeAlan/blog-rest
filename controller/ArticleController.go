package controller

import (
	"blog-rest/constant"
	"blog-rest/model"
	"blog-rest/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// parse article parameters
type  ArticleVo struct{
	Title   string `json:"title"`
	Content string `json:"content"`
	Author string `json:"author"`
}


func AddArticles(ctx *gin.Context){
	fmt.Println("add articles")
	
	var articleJson ArticleVo


	if ctx.BindJSON(&articleJson) == nil {
		if resultArticle, err := model.AddArticle(articleJson.Title, articleJson.Content,articleJson.Title); err == nil {
			// If the article is created successfully, show success message
			utils.DealHttpResponse(ctx, gin.H{
				"status":   constant.BlogSuccess,
				"message": constant.StatusText(constant.BlogSuccess),
				"data": resultArticle})
		} else {
			// if there was an error while creating the article, abort with an error
			ctx.AbortWithStatus(http.StatusBadRequest)
		} 
	}
 

	
}

