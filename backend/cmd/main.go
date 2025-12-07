package main

import (
	"context"
	"log"

	"github.com/denizdoganx/product-app/common/app"
	"github.com/denizdoganx/product-app/common/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	configurationManager := app.NewConfigurationManager()
	databaseInstance := mysql.GetConnectionPool(ctx, configurationManager.MySqlConfig)

	_, err := databaseInstance.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS productapp DEFAULT CHARACTER SET utf8mb4;")
	if err != nil {
		log.Fatalf("Database create error: %v", err)
	}

	_, err = databaseInstance.ExecContext(ctx, "USE productapp;")
	if err != nil {
		log.Fatalf("USE database error: %v", err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS product (
		id BIGINT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		price FLOAT NOT NULL,
		discount FLOAT NOT NULL,
		store VARCHAR(255) NOT NULL
	);`

	_, err = databaseInstance.ExecContext(ctx, createTableQuery)
	if err != nil {
		log.Fatalf("Table create error: %v", err)
	}

	e := echo.New()
	e.Start("0.0.0.0:8080")
}
