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
    return func(c *gin.Context) {
        var input struct {
            OriginalURL string `json:"original_url"`
        }
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }

        hash := md5.Sum([]byte(input.OriginalURL + time.Now().String()))
        shortName := hex.EncodeToString(hash[:])[:8]

        var existingURL models.ShortenedURL
        if err := db.Where("name = ?", shortName).First(&existingURL).Error; err == nil {
            c.JSON(http.StatusConflict, gin.H{"error": "Short name already exists"})
            return
        }

        shortenedURL := models.ShortenedURL{
            ShortCode:        shortName,
            URL: input.OriginalURL,
        }

        if err := db.Create(&shortenedURL).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create shortened URL"})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "shortened_url": shortName,
            "original_url":  input.OriginalURL,
        })
    }
}