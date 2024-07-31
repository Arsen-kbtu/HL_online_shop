package main

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID             uint      `gorm:"primaryKey" json:"id" readonly:"true" example:"1"`
	Name           string    `json:"name" validate:"required" example:"John Doe"`
	Email          string    `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Address        string    `json:"address" example:"123 Main St"`
	RegistrationAt time.Time `json:"registrationAt" readonly:"true" example:"2023-07-20T15:04:05Z"`
	Role           string    `json:"role" validate:"required,oneof=admin client" example:"client"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.RegistrationAt = time.Now()
	return
}
func (User) TableName() string {
	return "users_shop"
}
