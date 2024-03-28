package models

type Order struct {
	GormModel
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	CustomerName string `gorm:"type:varchar(100)" json:"customer_name"`
	OrderedAt    string `gorm:"type:varchar(100)" json:"ordered_at"`
	UserId       uint
	User         *User
}
