package gorm_sqllite

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/common"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func (d *Database) TryGetUnderlyingDatabaseConnection() interface{} {
	return d.db
}

func (d *Database) NewDatabase(name string) common.Database {
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", name)), &gorm.Config{})
	if err != nil {
		return nil
	}

	d.db = db

	return d
}

func DatabaseDriver() *Database {

	db := &Database{}

	return db

}

func AsGorm(database common.Database) *gorm.DB {
	raw := database.TryGetUnderlyingDatabaseConnection()

	if g, ok := raw.(*gorm.DB); ok {
		return g
	}

	return nil
}
