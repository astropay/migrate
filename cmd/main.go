package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"
)

// parametros:
//
//     db = database where to connect
//     environment = prod, stage, test
//     user = db user
//     password = db password
//     operation = up/down

var (
	db             string
	dbHost         string
	dbEnv          string
	dbUser         string
	dbPassword     string
	command        string
	migrationsPath string
)

func main() {

	// check required arguments
	if printHelp, err := parseArguments(); err != nil {
		if printHelp {
			printUsage()
		} else {
			log.Fatal(err.Error())
		}

		return
	}

	// run migration command
	if err := runMigrate(db, dbHost, dbEnv, dbUser, dbPassword, command, migrationsPath); err != nil {
		log.Fatal(err.Error())
	} else {
		log.Println("OK!")
	}

}

func parseArguments() (printUsage bool, err error) {

	flag.StringVar(&db, "db", "", "database schema migrations to run")
	flag.StringVar(&dbHost, "dbhost", "", "database host to connect; if present, 'env' is ignored")
	flag.StringVar(&dbEnv, "env", "", "environment where connecting in order to run the migrations: prod, stage, test")
	flag.StringVar(&dbUser, "dbuser", "", "database user to run the migration scripts")
	flag.StringVar(&dbPassword, "dbpass", "", "database password to run the migration scripts")
	flag.StringVar(&command, "command", "", "command to perform with scripts: 'up', 'down'")
	flag.StringVar(&migrationsPath, "dir", "", "migrations location")

	flag.Parse()

	printUsage, err = checkArguments()

	return
}

func checkArguments() (printUsage bool, err error) {

	// check environment
	if dbHost == "" {
		dbEnv = strings.ToLower(dbEnv)
		switch dbEnv {
		case "test", "prod", "staging", "dev":
			break
		case "":
			return true, fmt.Errorf("invalid parameter [env] value: '%s'", dbEnv)
		default:
			return false, fmt.Errorf("invalid parameter [env] value: '%s'", dbEnv)
		}
	}

	if dbEnv == "" && dbHost == "" {
		return false, errors.New("must indicate [host or [env] parameters")
	}

	// database
	if db == "" {
		return false, errors.New("invalid argument db")
	}

	// check user and password
	if dbUser == "" || dbPassword == "" {
		return false, errors.New("parameters [dbuser] and [dbpass] cannot be empty")
	}

	// check operation
	command = strings.ToLower(command)
	switch command {
	case "up", "down":
		break
	case "":
		return true, fmt.Errorf("invalid parameter [command] value: '%s'", command)
	default:
		return false, fmt.Errorf("invalid parameter [command] value: '%s'", command)
	}

	return
}

func printUsage() {
	fmt.Print("AstroPay Migration Tool\n\n")
	flag.Usage()
	fmt.Print("\nexamples:\n")
	fmt.Print("\tmigrate -env=test -db=apc -dbuser=myuser -dbpass=Passw0rd -command=up\n")
	fmt.Print("\tmigrate -host=127.0.0.1:3306 -db=apc -dbuser=myuser -dbpass=Passw0rd -command=up\n")
}
