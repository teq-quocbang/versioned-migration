package migrations

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

func Up(db *gorm.DB) error {
	getDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get DB, erorr: %v", err)
	}

	driver, err := mysql.WithInstance(getDB, &mysql.Config{MigrationsTable: "migration"})
	if err != nil {
		return fmt.Errorf("failed to instance, error: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://./migration", "example", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	log.Println("Up sucessfully")
	return nil
}
