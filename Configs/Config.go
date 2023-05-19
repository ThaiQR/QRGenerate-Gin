package config

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// "DB_USER": "root",
//
//	"DB_PASS": "Ms1793",
//	"DB_BBLGW_IP": "10.10.11.84",
//	"DB_BBLGW_NAME": "bblgw",
//	"DB_EKYC_NAME": "ekyc",
//	"DB_PORT": "3306",
//	"CHAR_SET": "utf8"
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "10.10.11.84",
		Port:     3306,
		User:     "root",
		Password: "Ms1793",
		DBName:   "bblgw",
	}

	return &dbConfig
}

func DBUrl(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
