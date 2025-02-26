package models

import (
    "time"
)

type ShortenedURL struct {
    ID uint `gorm: "primaryKey"`
    URL string `gorm:"not null"`
    ShortCode string `gorm:"not null"`
    AccessCount int    `gorm:"default:0"`
    CreatedAt *time.Time
    UpdatedAt *time.Time
}