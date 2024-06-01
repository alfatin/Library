run:
	go run main.go

run_db_migrate_up:
	go run database/migrate/up/up.go

run_db_migrate_down:
	go run database/migrate/down/down.go

run_seed:
	psql -U postgres library < database/seed.sql

run_algo:
	go run algoritma/algoritma.go