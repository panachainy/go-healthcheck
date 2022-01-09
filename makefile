ENV=\
	CLIENT_ID=xxx \
	CLIENT_SECRET=xxxxxxxxxx

up:
	$(ENV) \
	go run main.go test.csv

build:
	$(ENV) \
	go build

build.machine:
	go build

# require air in your machine
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
