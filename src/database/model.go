package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/gorm"
)

var Main *gorm.DB //Main database connection

var PGX *pgxpool.Pool
