all: paging

clean:
	rm -f paging

paging:
	go build -o paging cmd/paging/main.go

lint:
	golangci-lint run ./...

docker-up:
	docker compose -f docker-compose.yml build
	docker compose -f docker-compose.yml up -d

docker-down:
	docker compose -f docker-compose.yml down