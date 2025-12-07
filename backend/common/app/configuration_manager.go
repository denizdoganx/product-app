package app

import (
	"time"

	"github.com/denizdoganx/product-app/common/mysql"
)

type ConfigurationManager struct {
	MySqlConfig mysql.Config
}

func NewConfigurationManager() *ConfigurationManager {
	mySqlConfig := getMySqlConfig()
	return &ConfigurationManager{
		MySqlConfig: mySqlConfig,
	}
}

func getMySqlConfig() mysql.Config {
	return mysql.Config{
		Host:                  "product-app-mysql",
		Port:                  "3306",
		Username:              "root",
		Password:              "denizalper..2023",
		DbName:                "productapp",
		MaxConnections:        10,
		MaxIdleConnections:    5,
		MaxConnectionIdleTime: 5 * time.Minute,
		Timeout:               "5s",
		ReadTimeout:           "5s",
		WriteTimeout:          "5s",
	}
}
