# pggoose

This is a custom build of https://github.com/pressly/goose for Postgres only with simple defaults:

- Load .env
- Use DATABASE_URL for DB connection
- Use ./migrations dir

```
Usage: pggoose [OPTIONS] COMMAND

  -dir string
        directory with migration files (default "./migrations")
  -skip-env
        Don't load .env file

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
    fix
```
