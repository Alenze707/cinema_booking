
@REM The Data Source Name
set DB_DSN=postgres://postgres:postgres123@localhost:5432/cinema_booking?sslmode=disable

migrate -path ops/db_migrations -database %DB_DSN% goto %1

