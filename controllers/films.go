package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snsvistunov/go-pg-redis-memc/models"
)

// GET /film/:title
// Get film description by title
func GetFilm(context *gin.Context) {
	// Check if record exist
	start := time.Now()
	id := context.Param("title")
	film := new(models.Film)
	if film, found := models.GetMemcacheValue(models.MC, id); found {
		context.JSON(http.StatusOK, gin.H{"film": film})
		models.LogTracker(float32(time.Since(start).Microseconds())/(1000.00), "/film/"+id, "memory")
		return
	}
	if err := models.GetRedisValue(models.RC, id, film); err == nil {
		context.JSON(http.StatusOK, gin.H{"film": film})
		models.LogTracker(float32(time.Since(start).Microseconds())/(1000.00), "/film/"+id, "redis")
		return
	}

	if err := models.DB.Table("film").Where("title = ?", id).First(&film).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"film": film})
	models.LogTracker(float32(time.Since(start).Microseconds())/(1000.00), "/film/"+id, "db")
	models.SetMemcacheValue(models.MC, id, film)
	models.SetRedisValue(models.RC, id, film)
}
