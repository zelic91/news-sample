DATABASE_URL := $(shell grep POSTGRES_URL= .env | sed 's/POSTGRES_URL=//')

init:
	go mod tidy
	go get -u github.com/spf13/cobra@latest
	go get -u github.com/spf13/viper
	cp env.example .env
	mv gitignore .gitignore
	chmod +x deployment/gen_cert.sh
	chmod +x deployment/deploy.sh
	cd deployment && sh gen_cert.sh

dep:
	brew install dbmate
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

migrate:
	dbmate -u $(DATABASE_URL) -d ./db/postgres/migrations -s ./db/postgres/schema.sql up
	
rollback:
	dbmate -u $(DATABASE_URL) -d ./db/postgres/migrations -s ./db/postgres/schema.sql down

apigen:
	rm -rf api/gen
	mkdir api/gen
	oapi-codegen -package gen api/openapi.yaml > api/gen/spec.gen.go

dbgen_mongo:

dbgen_postgres:
	rm -rf db/postgres/dbgen
	sqlc -f db/postgres/sqlc.config.yaml generate

gen: migrate apigen dbgen_mongo dbgen_postgres
	go mod tidy
crawler:
	go run main.go crawler
server:
	go run main.go server
websocket:
	go run main.go websocket

docker:
	docker image rm news & \
	docker build -t news . && \
	docker run --rm \
		--env-file ./.env.docker \
		-p 3000:3000 \
		--name news news

test_api:
	curl localhost:3000/api/v1/auth/signup -X POST -H 'Content-Type: application/json' -d '{"username":"tester", "password":"123123123", "password_confirmation":"123123123"}'
	curl localhost:3000/api/v1/auth/signin -X POST -H 'Content-Type: application/json' -d '{"username":"tester", "password":"123123123"}'

.PHONY: init dev prepdev apigen toolsup toolsdown gen docker dbgen_mongo dbgen_postgres