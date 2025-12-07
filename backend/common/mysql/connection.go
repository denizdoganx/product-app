package mysql

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnectionPool(context context.Context, config Config) *sql.DB {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&timeout=%s&readTimeout=%s&writeTimeout=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
		config.Timeout,
		config.ReadTimeout,
		config.WriteTimeout,
	)

	database, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err)
	}

	database.SetMaxOpenConns(config.MaxConnections)
	database.SetMaxIdleConns(config.MaxIdleConnections)
	database.SetConnMaxIdleTime(config.MaxConnectionIdleTime)

	if err := database.PingContext(context); err != nil {
		panic(fmt.Sprintf("Unable to connect to MySQL: %v", err))
	}

	return database
}
