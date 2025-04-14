package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/cristianortiz/books-grpc/internal/pkg/configs"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

const (
	cutSet       = "file://" //represents the prefix to cut from migration directory path
	databaseName = "postgres"
)

type Migrator struct {
	pgDBMigrate *migrate.Migrate
}

// provides a migrator instance, it takes a db config and gorm dbConn as input. It initializes and returns
// a migrator instance with the postgres db migrator
func ProvideMigrator(config configs.DatabaseConfig, pgDB *gorm.DB) (*Migrator, error) {
	dbConn, err := pgDB.DB()
	if err != nil {
		return nil, err
	}

	pgDBMigrate, err := initMigrate(dbConn, config.MigrationPath)
	if err != nil {
		return nil, err
	}
	return &Migrator{
		pgDBMigrate: pgDBMigrate,
	}, nil
}

func (m Migrator) RunMigrations() {
	m.RunMigrationsWith(m.pgDBMigrate, "Postgres Database")
}
func (m Migrator) RunMigrationsWith(migrateInstance *migrate.Migrate, dbName string) {
	if err := migrateInstance.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Printf("no change detected after running the migrations for %s", dbName)
			return
		}
		log.Println(fmt.Sprintf("migration Failed for %s", dbName), err)
	}
	log.Printf("migrations applied succesfully to  %s", dbName)

}

// initializes the migration instance for postgres db. it takes sql.DB connection and a migrations
// directory as input. It creates a postgres driver instance and then initializes the migration inst
// using db driver
func initMigrate(dbConn *sql.DB, dir string) (*migrate.Migrate, error) {
	//creates postgres driver instance
	driver, err := postgres.WithInstance(dbConn, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	sourcePath, err := getSourcePath(dir)
	if err != nil {
		return nil, err
	}
	//initializes migrations instance usgin db driver
	m, err := migrate.NewWithDatabaseInstance(sourcePath, databaseName, driver)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// process the migration directory path. It trims the "file://" prefix, obtains the absolute path,
// and returns the formattes source path required for migration
func getSourcePath(dir string) (string, error) {
	dir = strings.TrimPrefix(dir, cutSet)
	// absolute path is important to avoid error finding the migrations directory route, ex, if the route is defined as relative
	// or outside the project directory or in other level, ex input "migrations" (relative ) and app executes
	//  from /app, filepath.abs returns "/app/migrations" (absolute)
	absPath, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}
	//now add again prefix required for migrate pkg to the abs path: returns "file:///app/migrations"
	return fmt.Sprintf("%s%s", cutSet, absPath), nil
}
