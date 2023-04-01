package models

import (
	"time"
	// "gorm.io/gorm"
)

type Activity struct {
	Id        int64     `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type Todo struct {
	Id              int64     `json:"id" gorm:"primary_key"`
	ActivityGroupId int64     `json:"activity_group_id""`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	Title           string    `json:"title"`
	UpdatedAt       time.Time `json:"updatedAt"`
	CreatedAt       time.Time `json:"createdAt"`
}
