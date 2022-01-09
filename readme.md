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

2. Run app in watch mode

    `make up.watch`

    > require air

3. (alternative) Run app

    `make up`

#### Test

`make test`

### Usage

1. Build

    | Command            | Description                |
    | ------------------ | -------------------------- |
    | make build         | build with env in makefile |
    | make build.machine | build with env in machine  |

2. Run `go-healthcheck test.csv`

## CSV

https://docs.google.com/spreadsheets/d/1Wl0Q9OPbjx1j9FbwXbrfEvtb9wzw5HHZ1mZIDIVN8zs/edit?usp=sharing

## Note

### Handle errors without stopping the entire process

I'm not sure about this requirement, So I think some process if it error should be stop process because this project is cli.

Example: If you can't connect to line api for get token then you can't submit summary right?

Example2: If you can't connect to health check some url or timeout should be handle it. (in this case I think the same as you)

## REF

- [line-login-sdk](https://www.youtube.com/watch?v=dimWmt2RHiU)
- [line-login-easy](https://jaedsada.me/blogs/blog/line-oauth)
- [cli-oauth](https://gist.github.com/marians/3b55318106df0e4e648158f1ffb43d38)
- [csv-reader](https://golang.cafe/blog/golang-read-file-example.html)
