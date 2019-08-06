package main

import (
	"fmt"
	"os"
	"path"

	"github.com/astropay/go-tools/files"
	"github.com/pressly/goose"

	_ "github.com/go-sql-driver/mysql"
)

// main migration logic
func runMigrate(db, env, dbUser, dbPass, command, migrationsDir string) (err error) {

	currentPath, _ := os.Getwd()
	migrationsPath := ""

	if migrationsDir == "" {
		migrationsPath = path.Clean(currentPath)
	} else {
		migrationsPath = migrationsDir
	}

	// check if schema directory exists
	schemaPath := path.Join(migrationsPath, db)
	if !files.Exists(schemaPath) {
		return fmt.Errorf("schema directory '%s' does not exist", schemaPath)
	}

	// build conn string
	host := getHost(env)
	connStr := fmt.Sprintf(ConnectionString, dbUser, dbPass, host, db)

	// open db connection
	dbDriver, errOpenDB := goose.OpenDBWithDriver(MySQLDriver, connStr)
	if errOpenDB != nil {
		return fmt.Errorf("goose: failed to open DB @ %s: %s", host, errOpenDB)
	}

	// run
	if errRun := goose.Run(command, dbDriver, schemaPath); errRun != nil {
		return fmt.Errorf("goose: %s", errRun)
	}

	return
}

func getHost(env string) (host string) {
	switch env {
	case EnvProduction:
		host = DBHostProduction
	case EnvStaging:
		host = DBHostStaging
	case EnvTesting:
		host = DBHostTesting
	case EnvDevelopment:
		host = DBHostDev
	}

	return
}
