package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/alifnh/bjb-auction-backend/internal/config"
	"github.com/alifnh/bjb-auction-backend/internal/pkg/logger"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB(cfg *config.Config) (*sql.DB, error) {
	dbCfg := cfg.Database

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbCfg.Host,
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.DbName,
		dbCfg.Port,
		dbCfg.Sslmode,
	)

	db, err := sql.Open("pgx", dsn)

	if err != nil {
		logger.Log.Fatal("error initializing database: ", err.Error())
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(dbCfg.MaxIdleConn)
	db.SetMaxOpenConns(dbCfg.MaxOpenConn)
	db.SetConnMaxLifetime(time.Duration(dbCfg.MaxConnLifetimeMinute) * time.Minute)

	return db, nil
}
