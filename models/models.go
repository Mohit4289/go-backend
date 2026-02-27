package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username string    `gorm:"size:20;unique;not null"`
	Password string    `gorm:"size:500;not null"`
	Email    string    `gorm:"size:100;unique;not null"`
}

type Todo struct {
	ID     uint      `gorm:"primaryKey"`
	Text   string    `gorm:"not null"`
	Status string    `gorm:"not null;default:'pending'"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`
}

func (User) TableName() string {
	return "users"
}

func (Todo) TableName() string {
	return "todos"
}
