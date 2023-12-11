build:
	docker-compose build

run:
	docker-compose up -d

stop:
	docker-compose down

test:
	go test -cover ./...

lint:
	golangci-lint run -v


