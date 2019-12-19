package main

import (
	"GinHello/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run("0.0.0.0:8080")
}
