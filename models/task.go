package models

import "time"

type Task struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Title     string     `gorm:"not null" json:"title"`
	Completed bool       `gorm:"default:false" json:"completed"`
	Priority  string     `gorm:"type:varchar(10);default:'medium'" json:"priority"` // low, medium, high
	DueDate   *time.Time `json:"due_date,omitempty"` // nullable
	UserID    uint       `gorm:"not null" json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
}