package main

import (
	"textile_flannel/internal/database"
	"textile_flannel/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
    db := database.InitDB()
	sqlDb, _ := db.DB()

    defer sqlDb.Close()

    r := gin.Default()

    r.POST("/shorten", handlers.ShortenURL(db))
    r.GET("/shorten/:name", handlers.ShowURL(db))
    r.PUT("/shorten/:name", handlers.UpdateURL(db))
    r.DELETE("/shorten/:name", handlers.DeleteURL(db))
    r.GET("/shorten/:name/stats", handlers.GetStats(db))

    r.Run(":8080")
}