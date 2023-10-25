package connection

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConfig struct {
	Address  string
	Port     int
	Name     string
	Username string
	Password string
}

func NewDatabaseConn(m MySQLConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Address, m.Port, m.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
