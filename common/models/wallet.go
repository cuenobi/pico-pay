package models

type Wallet struct {
	BaseModel

	Code   string  `gorm:"unique" json:"code"`
	Name   string  `gorm:"default:'your wallet'" json:"name"`
	Amount float32 `gorm:"default:0" json:"amount"`
}
