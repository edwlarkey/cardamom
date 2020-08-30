package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/edwlarkey/cardamom/pkg/models"
)

type DB struct {
	DB *gorm.DB
}

func (d *DB) Connect(db_type string, dsn string) error {
	db, err := gorm.Open(db_type, dsn)
	if err != nil {
		return err
	}

	d.DB = db
	return nil
}

func (d *DB) Migrate() error {
	d.DB.LogMode(true)
	if err := d.DB.AutoMigrate(&models.User{}, &models.Bookmark{}, &models.Tag{}).Error; err != nil {
		return err
	}

	return nil
}

func (d *DB) Close() {
	d.DB.Close()
}
