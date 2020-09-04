package db

import (
	"github.com/edwlarkey/cardamom/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func (d *DB) Connect(dialect string, dsn string) error {
	var db *gorm.DB
	var err error
	switch dialect {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}
	}

	d.DB = db
	return nil
}

func (d *DB) Migrate() error {
	err := d.DB.AutoMigrate(&models.User{}, &models.Bookmark{}, &models.Tag{})
	if err != nil {
		return err
	}

	return nil
}
