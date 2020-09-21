package models

import (
	"gorm.io/gorm"
	"time"
)

type (
	User struct {
		gorm.Model
		Id        string `gorm:"primarykey"`
		Name      *string
		Email     *string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}
)
