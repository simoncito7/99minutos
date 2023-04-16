postgres:
	# To start a new postgres container.
	docker run --name logistic-app -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:15-alpine
createdb:
	# To create a new database called logistic_app inside the container *logistic-app*.
	docker exec -it logistic-app createdb --username=root --owner=root logistic_app
dropdb:
	# To drop the database.
	docker exec -it logistic-app dropdb logistic_app
createmigration:
	# To create the files with the SQL statements (migrate up and down).
	migrate create -ext sql -dir db/migrations -seq init_schema
migrateup:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/logistic_app?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/logistic_app?sslmode=disable" -verbose down
test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test