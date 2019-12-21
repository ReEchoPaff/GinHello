package initRouter

import (
	"GinHello/handler"
	"GinHello/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
)

func SetupRouter() *gin.Engine {
	rootPath, _ := os.Getwd()
	router := gin.New()
	// 添加自定义的logger中间件， context.Next()之前顺序执行，之后逆序执行，类似于栈。
	router.Use(gin.Logger(), middleware.Logger(), gin.Recovery())

	if mode := gin.Mode(); mode == gin.TestMode || mode == gin.DebugMode {
		router.LoadHTMLGlob("./templates/*")
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
		// 根据id删除一篇文章
		articleRouter.DELETE("/article/:id", handler.DeleteOne)
		// 根据id获取一篇文章
		articleRouter.GET("/article/:id", handler.GetOne)
		// 获取所有的文章
		articleRouter.GET("/articles", handler.GetAll)
	}

	// 添加swagger
	url := ginSwagger.URL("http://192.168.121.134:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
