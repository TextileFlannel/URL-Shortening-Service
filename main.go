package main

import (
	"textile_flannel/internal/config"
	"textile_flannel/internal/database"
	"textile_flannel/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
    cfg := config.LoadConfig()

    db := database.InitDB()
	sqlDb, _ := db.DB()

    defer sqlDb.Close()

    r := gin.Default()

    r.POST("/shorten", handlers.ShortenURL(db))
    //r.GET("/shorten/:name", internal.RedirectURL(db))
    //r.PUT("/shorten/:name", internal.UpdateURL(db))
    //r.DELETE("/shorten/:name", internal.DeleteURL(db))
    //r.GET("/shorten/:name/stats", internal.GetStats(db))

    r.Run(cfg.ServerPort)
}