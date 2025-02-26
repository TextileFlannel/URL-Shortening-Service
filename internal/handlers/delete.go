package handlers

import (
	"net/http"
	"textile_flannel/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteURL(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")

		res := db.Where("short_code = ?", name).Delete(&models.ShortenedURL{})
		if res.Error != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		
		ctx.AbortWithStatus(http.StatusNoContent)
	}
}