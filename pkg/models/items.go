package models

type Item struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ItemCode    string `gorm:"type:varchar(100)" json:"item_code"`
	Description string `gorm:"type:varchar(200)" json:"description"`
	Quantity    uint   `gorm:"type:int" json:"quantity"`
}
