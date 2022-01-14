.PHONY: all run

all: prestart run

prestart:
	docker-compose up -d

run:
	go run ./cmd/main.go

down:
	docker-compose down --remove-orphans
