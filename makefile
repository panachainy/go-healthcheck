ENV=\

up:
	$(ENV) \
	go run main.go

build:
	$(ENV) \
	go build

# require air in your machine
up.watch:
	$(ENV) \
	air
