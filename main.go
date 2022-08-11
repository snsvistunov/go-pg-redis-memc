package main

import (
	"github.com/gin-gonic/gin"
	"github.com/snsvistunov/go-pg-redis-memc/controllers"
	"github.com/snsvistunov/go-pg-redis-memc/models"
)

func main() {
	//New Gin router
	router := gin.Default()

	//New pg connection
	models.ConnectDB()
	
	//New Redis connection
	models.ConnectRedis()
	
	//New app in-memory cache
	models.ConnectMemcache()
	
	//Router handlers assertion
	router.GET("/film/:title", controllers.GetFilm)
	router.GET("/log", controllers.GetTimeLog)
	
	//Start router
	router.Run()

}
