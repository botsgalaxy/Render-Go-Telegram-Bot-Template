package database

import "gorm.io/gorm"

type TelegramUser struct {
	gorm.Model
	UserId       int64 `gorm:"unique"`
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	IsPremium    bool
}

func (t *TelegramUser) Create() error {
	result := DB.Create(t)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
