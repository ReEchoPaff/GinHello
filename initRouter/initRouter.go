package initRouter

import (
	"GinHello/handler"
	"GinHello/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func SetupRouter() *gin.Engine {
	rootPath, _ := os.Getwd()
	router := gin.New()
	// 添加自定义的logger中间件， context.Next()之前顺序执行，之后逆序执行，类似于栈。
	router.Use(gin.Logger(), middleware.Logger(), gin.Recovery())

	if mode := gin.Mode(); mode == gin.TestMode || mode == gin.DebugMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.StaticFS("/avatar", http.Dir(rootPath+"/avatar/"))
	router.Static("/statics", "./statics/")
	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}

	// 添加 user
	userRouter := router.Group("/user")
	{
		// 注册
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/rest_register", handler.RestUserRegister)
		// 登陆
		userRouter.POST("/login", handler.UserLogin)
		userRouter.POST("/rest_login", handler.RestUserLogin)
		// 通过id，返回用户信息
		userRouter.GET("/profile/", middleware.Auth(), handler.UserProfile)
		userRouter.GET("/rest_profile", middleware.Auth(), handler.RestUserProfile)
		// 上传文件
		userRouter.POST("/update", middleware.Auth(), handler.UpdateUserProfile)
	}

	// 添加 article
	articleRouter := router.Group("")
	{
		// 添加一篇文章
		articleRouter.POST("/article", handler.ArticleInsert)
	}

	return router
}
