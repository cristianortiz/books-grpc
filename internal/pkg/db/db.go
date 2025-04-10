package db

import (
	"fmt"
	"log"
	"time"

	"github.com/cristianortiz/books-grpc/internal/pkg/configs"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DatabasePool interface {
	SetMaxOpenConns(n int)
	SetConnMaxLifetime(d time.Duration)
	SetMaxIdleConns(n int)
	SetConnMaxIdleTime(d time.Duration)
}

func ProvideDBConn(config *configs.DatabaseConfig) (*gorm.DB, error) {
	dbName := config.Dbname
	username := config.Username
	password := config.Password
	host := config.Host
	port := config.Port
	sslMode := config.SslMode
	connectionTimeOut := config.Connection.TimeOut

	args := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s connect_timeout=%d", host, port, username, dbName, password, sslMode, connectionTimeOut)
	dbConn, err := gorm.Open(postgres.Open(args), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: fmt.Sprintf("%s.", config.Schema),
		},
	})

	if err != nil {
		return nil, errors.WithMessage(err, "getDB.gorm.Open: failed to connect to db")
	}
	sqlDB, err := dbConn.DB()
	if err != nil {
		return nil, errors.WithMessage(err, "Setup.db.DB")
	}
	ConfigureDatabasePool(sqlDB, config.Connection)
	if err := sqlDB.Ping(); err != nil {
		return nil, errors.WithMessage(err, "setup.sqlDB.Ping")
	}
	log.Println("created DB connection pool", sqlDB.Stats().OpenConnections, sqlDB.Stats().MaxOpenConnections, sqlDB.Stats().InUse, sqlDB.Stats().Idle)
	return dbConn, nil
}

func ConfigureDatabasePool(pool DatabasePool, conn configs.ConnectionPool) {

	pool.SetMaxOpenConns(conn.MaxOpenConnections)
	pool.SetConnMaxLifetime(time.Second * time.Duration(conn.MaxLifeTime))
	pool.SetMaxIdleConns(conn.MaxIdleConnections)
	pool.SetConnMaxIdleTime(time.Second * time.Duration(conn.MaxIdleTime))
}
