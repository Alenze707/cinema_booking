## DB Notes

To create a new database migration files (the pair of up and down files), use such example:

PS D:\Aplicatii\Go-Projects\cinema_booking> migrate create -seq -ext.sql -dir=ops/db_migrations rooms
D:\Aplicatii\Go-Projects\cinema_booking\ops\db_migrations\000001_rooms.up.sql
D:\Aplicatii\Go-Projects\cinema_booking\ops\db_migrations\000001_rooms.down.sql
PS D:\Aplicatii\Go-Projects\cinema_booking>

