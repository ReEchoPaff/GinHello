package main

import (
	_ "GinHello/docs"
	"GinHello/initRouter"
)

// @title Gin swagger hello
// @version 1.1
// @description Gin swagger 示例项目

// @contact.name repaff
// @contact.url http://bing.com
// @contact.email guozerun33@163.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 192.168.121.134:8080
func main() {
	router := initRouter.SetupRouter()
	_ = router.Run("192.168.121.134:8080")
}
