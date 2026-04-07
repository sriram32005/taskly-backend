package models

import "time"

type Task struct {
	ID        uint       `gorm:"primaryKey"`
	Title     string     `gorm:"not null"`
	Completed bool       `gorm:"default:false"`
	Priority  string     `gorm:"type:varchar(10);default:'medium'"` // low, medium, high
	DueDate   *time.Time // nullable
	UserID    uint       `gorm:"not null"`
	CreatedAt time.Time
}