up:
	go run main.go test.csv

build:
	go build

start:
	go-healthcheck test.csv

# 1. require air in your machine
# 2. you must comment // Receive csvPath from argument. section to use hardcode instead for use air to develop.
#    example hardcode -> csvPath := "test.csv"
up.watch:
	air

tidy:
	go mod tidy

mock:
	cd mocks && npm run mock

mock.down:
	cd mocks && npm run mock:down

test:
	go test -v -cover ./...

test.cov:
	go test -v -race -covermode=atomic -coverprofile=covprofile ./...

# Docker have issue (follow detail in readme.md)
docker.build:
	docker build -t go-healthcheck .

docker.run:
	docker run --rm -p 14565:14565 go-healthcheck test.csv
