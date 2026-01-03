package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"github.com/joho/godotenv"
)

var (
	flags   = flag.NewFlagSet("pggoose", flag.ExitOnError)
	dir     = flags.String("dir", "./migrations", "directory with migration files")
	skipEnv = flags.Bool("skip-env", false, "Don't load .env file")

	usagePrefix = `
This is a custom build of https://github.com/pressly/goose for Postgres only
With simple defaults:
* Load .env
* Use DATABASE_URL for DB connection
* Use ./migrations dir

Usage: pggoose [OPTIONS] COMMAND
	`
	usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-by-one            Migrate the DB up by 1
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    reset                Roll back all migrations
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with the current timestamp
    fix                  Apply sequential ordering to migrations
`
)

func main() {
	flags.Usage = usage

	if len(os.Args) < 2 {
		flags.Usage()
		return
	}

	if !*skipEnv {
		godotenv.Load()
	}

	command := os.Args[1]
	args := os.Args[2:]

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := goose.OpenDBWithDriver("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	if err := goose.RunContext(ctx, command, db, *dir, args...); err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Println(usagePrefix)
	flags.PrintDefaults()
	fmt.Println(usageCommands)
}
