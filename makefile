ENV=\
	CLIENT_ID=xxx \
	CLIENT_SECRET=xxxxxxxxxx

up:
	$(ENV) \
	go run main.go test.csv

build:
	$(ENV) \
	go build

# require air in your machine
up.watch:
	$(ENV) \
	air

tidy:
	go mod tidy
