ENV=\
	CLIENT_ID=xxx \
	CLIENT_SECRET=xxxxxxxxxx

up:
	$(ENV) \
	go run main.go test.csv

build:
	go build

start:
	$(ENV) \
	go-healthcheck test.csv

# require air in your machine
# you must comment // Receive csvPath from argument. section to use hardcode instead for use air to develop.
# example hardcode -> csvPath := "test.csv"
up.watch:
	$(ENV) \
	air

tidy:
	go mod tidy

mock:
	cd mocks && npm run mock

mock.down:
	cd mocks && npm run mock:down

test:
	go test -v -cover ./...

# Docker have issue (follow detail in readme.md)
docker.build:
	docker build -t go-healthcheck .

docker.run:
	docker run --rm -p 14565:14565 go-healthcheck test.csv
