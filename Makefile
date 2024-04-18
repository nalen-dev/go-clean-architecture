migrateup:
	migrate -path internal/database/migration -database "postgresql://postgres:password@localhost/be_test" -verbose up

migratedown:
	migrate -path internal/database/migration -database "postgresql://postgres:password@localhost/be_test" -verbose down

migratefix:
	migrate -path internal/database/migration -database "postgresql://postgres:password@localhost/be_test" force 1