package db

import (
	"github.com/edwlarkey/cardamom/pkg/models"
	"gorm.io/gorm"
)

func currentUser(user *models.User) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id", user.ID)
	}
}
