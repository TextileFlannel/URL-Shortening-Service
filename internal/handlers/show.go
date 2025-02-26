package handlers

import (
	"net/http"
	"textile_flannel/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShowURL(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")

		var data models.ShortenedURL
		if err := db.Where("short_code = ?", name).First(&data).Error; err != nil{
			ctx.JSON(http.StatusNotFound, gin.H{"error": "url not found"})
			return
		}
		data.AccessCount += 1
		db.Save(&data)
		
		ctx.JSON(http.StatusOK, gin.H{
			"id": data.ID,
			"url": data.URL,
			"shortCode": data.ShortCode,
			"createdAt": data.CreatedAt,
			"updatedAt": data.UpdatedAt,
		})
	}
}