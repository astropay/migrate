package main

// Database server addresses
const (
	DBHostProduction = "master.apc.db.astropay:3306"
	DBHostTesting    = "master.apc.db.astropay.test:3306"
	DBHostStaging    = "master.apc.db.astropay.stg:3306"
	DBHostDev        = "127.0.0.1:3306"
)

// Goose related constants
const (
	MySQLDriver      = "mysql"
	ConnectionString = "%s:%s@tcp(%s)/%s?parseTime=true" // "{%dbuser}:{%dbpass}@tcp({%host})/{%schema}?parseTime=true"
)

// Environments
const (
	EnvProduction  = "prod"
	EnvStaging     = "stage"
	EnvTesting     = "test"
	EnvDevelopment = "dev"
)
