package db

import (
	"errors"

	"github.com/edwlarkey/cardamom/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DialectNotSupported = errors.New("Dialect not supported. [sqlite, postgres]")

type DB struct {
	DB *gorm.DB
}

func (d *DB) Connect(dialect string, dsn string) error {
	var db *gorm.DB
	var err error
	switch dialect {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
		if err != nil {
			return err
		}
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}
	default:
		return DialectNotSupported
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
