package handler

import (
	"GinHello/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArticleInsert(context *gin.Context) {
	article := model.Article{}
	var id = -1
	if e := context.ShouldBindJSON(&article); e == nil {
		id = article.Insert()
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"id":   id,
	})
}
