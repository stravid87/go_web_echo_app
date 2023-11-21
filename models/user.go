package models

type User struct {
	ID       uint   `gorm:"primaryKey; unique; autoIncrement; not null" json:"id"`
	Email    string `gorm:"column:email; unique; not null" json:"email"`
	Username string `gorm:"column:username; unique; not null" json:"username"`
	Password string `gorm:"column:password; not null" json:"password"`
}

func (User) TableName() string {
	return "users"
}
