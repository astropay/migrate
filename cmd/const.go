package main

// Database server addresses
const (
	DBHostProduction = "master.apc.db.prod.astro:3306"
	DBHostTesting    = "apc.db.tst.astro:3306"
	DBHostStaging    = "master.apc.db.stg.astro:3306"
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
