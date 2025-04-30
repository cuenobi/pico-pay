package models

type User struct {
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `gorm:"not null" json:"password"`

	WalletID *string      `gorm:"not null;uniqueIndex:idx_wallet" json:"wallet_id"`
	Wallet   *Wallet `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"wallet"`

	BaseModel
}
