package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snsvistunov/go-pg-redis-memc/models"
)

// GET /logs
// Get film description by title
func GetTimeLog(context *gin.Context) {
	var logs []models.ResponseTimeLog
	result := models.DB.Find(&logs)
	context.JSON(http.StatusOK, gin.H{"ResponseTimeLog": result.Value})
}
