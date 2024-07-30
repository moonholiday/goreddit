.PHONY = postgres adminer migrate

postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} postgres

adminer:
	docker run --rm -ti --network host adminer

migrate:
	migrate -source file://migrations -database postgres://postgres:${POSTGRES_PASSWORD}@localhost/postgres?sslmode=disable up

