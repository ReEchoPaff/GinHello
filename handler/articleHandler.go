package handler

import (
	"GinHello/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// @Summary 插入一篇文章
// @Tags 文章
// @version 1.0
// @Accept application/x-json-stream
// @Param article body model.Article true "文章"
// @Success 200 object model.Result 成功后返回的值
// @Failure 409 object model.Result 失败
// @Router /article [post]
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

// @Summary 根据id获取一篇文章
// @Tags 文章
// @version 1.0
// @Accept application/x-json-stream
// @Param id path int true "id"
// @Success 200 object model.Result 成功后返回的值
// @Router /article/{id} [get]
func GetOne(context *gin.Context) {
	id := context.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		log.Panicln("id 不是int类型，id转换失败", e.Error())
	}
	article := model.Article{
		ID: i,
	}
	art := article.FindById()
	context.JSON(http.StatusOK, gin.H{
		"article": art,
	})
}

// 获取所有的文章
func GetAll(context *gin.Context) {
	article := model.Article{}
	articles := article.FindAll()
	context.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}

// 根据id删除对应的文章
func DeleteOne(context *gin.Context) {
	id := context.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		log.Panicln("id 不是 int类型， id转换失败", e.Error())
	}
	article := model.Article{ID: i}
	article.DeleteOne()
}
