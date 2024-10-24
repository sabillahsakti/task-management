package models

import "time"

type Task struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Status      string    `gorm:"type:varchar(50);default:'pending'" json:"status"`
	UserID      int64     `gorm:"not null" json:"user_id"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DueDate     time.Time `json:"due_date"`
	Priority    string    `gorm:"type:varchar(50);default:'medium'" json:"priority"`
	User        User      `gorm:"foreignKey:UserID"`
}
