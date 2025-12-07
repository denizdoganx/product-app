package infrastructure

import (
	"context"
	"database/sql"

	"github.com/labstack/gommon/log"
)

func TruncateTestData(ctx context.Context, dbPool *sql.DB) {
	_, truncateResultErr := dbPool.ExecContext(ctx, "TRUNCATE TABLE products")
	if truncateResultErr != nil {
		log.Error(truncateResultErr)
	} else {
		log.Info("Products table truncated")
	}
}
