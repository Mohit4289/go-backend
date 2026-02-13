package models

type User struct {
	ID       string `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Email    string `gorm:"column:email"`
}

func (User) TableName() string {
	return "users"
}
