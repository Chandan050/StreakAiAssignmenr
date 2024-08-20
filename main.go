package main

import (
	"myapp/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
  _ = Router()
}

func Router() *gin.Engine {
		// Create a new instance of the logger
		router := gin.Default()

		//router for Find-pair method
		router.POST("/find-pairs", controllers.FindPairAPi)
	
		//runs at default port 8080
		router.Run()
		return router
}