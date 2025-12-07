package infrastructure

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/labstack/gommon/log"
)

var INSERT_PRODUCTS = `INSERT INTO products (name, price, discount,store) VALUES('AirFryer',3000, 22, 'ABC TECH'),
('Ütü',1500, 10, 'ABC TECH'),
('Çamaşır Makinesi',10000, 15, 'ABC TECH'),
('Lambader',2000, 0, 'Dekorasyon Sarayı');`

func TestDataInitialize(ctx context.Context, dbPool *sql.DB) {
	insertProductsResult, insertProductsErr := dbPool.ExecContext(ctx, INSERT_PRODUCTS)
	if insertProductsErr != nil {
		log.Error(insertProductsErr)
	} else {
		affectedRowCount, _ := insertProductsResult.RowsAffected()

		log.Info(fmt.Sprintf("Products data created with %d rows", affectedRowCount))
	}
}
