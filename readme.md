## Cinema Booking

<br/>

### DB Notes

As prerequisites, install the following tools:
1. [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) tool for running the database migrations.
2. [air](https://github.com/cosmtrek/air) tool used in `run_dev` to have the app restarted on code changes.

Use `run_db_migrate` script (`.bat` is just Windows specific). To provide the proper version,
checkout the latest version in `ops/db_migrations` directory.<br/>
Then run it like this (as an example): `run_db_migrate 1`

<br/>
