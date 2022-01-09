# Go Health check

## Prerequisites

- go version go1.17 darwin/amd64
- Mockoon
  - require node v14.17.6
- air
- docker [optional]

### Setup

#### Mockoon

- macOS: Use Mockoon via apps install -> `brew install --cask mockoon`

#### Env

##### Makefile (for only develop)

Update file `./makefile` to your clientId(Channel ID), SecretId(Channel secret)

> I added both of keys in file secret.txt root of zip file.

```sh
# ./makefile

ENV=\
  CLIENT_ID=0000000000 \
  CLIENT_SECRET=11111111111111111111111111111111

```

##### Machine environment (for macOS usage)

Run script below (change to your key instead)

```sh
export CLIENT_ID=0000000000
export CLIENT_SECRET=11111111111111111111111111111111
```

> You can add in `vi ~/.bash_profile` up to your machine.

### Development

1. Start Mockoon for prepair external service mock.

    `make mock`

2. Run app

    `make up`

3. (alternative) Run app in watch mode

    `make up.watch`

    > require air
    >
    > require mock

#### Test

`make test`

#### Create class mock

example `mockgen -source=healthz/externals/health.go -destination=healthz/externals/mock_externals/health.go`

### Usage

1. Build `make build`
2. Run `go-healthcheck test.csv`

### CLI

| Command        | Description                                     |
| -------------- | ----------------------------------------------- |
| make up        | run app with basic command                      |
| make build     | build app to get artifact file `go-healthcheck` |
| make start     | run app from artifact file `go-healthcheck`     |
| make tidy      | clean up go.mod go.sum                          |
| make mock      | run mock in background (port 9091)              |
| make mock.down | down mock                                       |
| make test      | run test (require mock)                         |

## CSV

https://docs.google.com/spreadsheets/d/1Wl0Q9OPbjx1j9FbwXbrfEvtb9wzw5HHZ1mZIDIVN8zs/edit?usp=sharing

## Note

1. Handle errors without stopping the entire process

    I'm not sure about this requirement, So I think some process if it error should be stop process because this project is cli.

    Example: If you can't connect to line api for get token then you can't submit summary right?

    Example2: If you can't connect to health check some url or timeout should be handle it. (in this case I think the same as you)

2. Have issue when make app in CLI way
   1. Can't ref csv file outside dockerfile
   2. Can't open web for login oauth)

    > That my mistake for choose make it in CLI way.

## REF

- [line-login-sdk](https://www.youtube.com/watch?v=dimWmt2RHiU)
- [line-login-easy](https://jaedsada.me/blogs/blog/line-oauth)
- [cli-oauth](https://gist.github.com/marians/3b55318106df0e4e648158f1ffb43d38)
- [csv-reader](https://golang.cafe/blog/golang-read-file-example.html)
- [external-mock](https://wawand.co/blog/posts/how-to-mock-an-external-service-for-test-in-go-73251a7a/)
