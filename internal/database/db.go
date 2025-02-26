package database

import (
    "textile_flannel/internal/models"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

func InitDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&models.ShortenedURL{})
    return db
}