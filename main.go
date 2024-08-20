package main

import (
	"myapp/controllers"

	"github.com/gin-gonic/gin"
)

//title Documentation
//@version 1.0
//@Description this is sample API for finding sum of indics equal the given target
//@host localhost:8080
//BasePath /
func main() {
  _ = Router()
}

// contains all Router 
func Router() *gin.Engine {
		// Create a new instance of the logger
		router := gin.Default()

		//router for Find-pair method
		router.POST("/find-pairs", controllers.FindPairAPi)
	
		//runs at default port 8080
		router.Run()
		return router
}