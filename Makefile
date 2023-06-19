postgres:
	sudo docker run -d --name postgres15alpl --network bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.3-alpine3.17
createdb:
	sudo docker exec -it postgres15alpl createdb --username=root --owner=root simplebank

dropdb:
	sudo docker exec -it postgres15alpl dropdb simplebank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" -verbose down 1


sqlc:
	sqlc generate

test:
	go test -v -cover ./db/sqlc


server:
	go run main.go


mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/dasotd/simplebank/db/sqlc Store

docker:
	sudo docker run --name simplebank --network bank -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres15alpl:5432/simplebank?sslmode=disable" simplebank:latest

migration:
	migrate create -ext sql -dir  db/migration -seq {theName}

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	proto/*.proto
	
evans:
	evans --host localhost --port 9090 -r repl


.PHONY: creatdb dropdb postgres migratedown migrateup migratedown1 migrateup1 test server mock docker migration db_docs db_schema proto evans

# migrate create -ext sql -dir db/migration -seq add_users
# statik -src=./doc/swagger -dest=./doc
# --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	# --openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \

# rm -f doc/swagger/*.swagger.json