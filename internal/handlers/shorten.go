package handlers

import (
    "crypto/md5"
    "encoding/hex"
    "net/http"
    "textile_flannel/internal/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "time"
)

func ShortenURL(db *gorm.DB) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var input struct {
            OriginalURL string `json:"url"`
        }
        if err := ctx.ShouldBindJSON(&input); err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }

        hash := md5.Sum([]byte(input.OriginalURL + time.Now().String()))
        shortName := hex.EncodeToString(hash[:])[:8]

        var existingURL models.ShortenedURL
        if err := db.Where("short_code = ?", shortName).First(&existingURL).Error; err == nil {
            ctx.JSON(http.StatusConflict, gin.H{"error": "Short name already exists"})
            return
        }

        shortenedURL := models.ShortenedURL{
            ShortCode:        shortName,
            URL: input.OriginalURL,
        }

        if err := db.Create(&shortenedURL).Error; err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create shortened URL"})
            return
        }

        ctx.JSON(http.StatusCreated, gin.H{
            "short_code": shortName,
            "url":  input.OriginalURL,
        })
    }
}