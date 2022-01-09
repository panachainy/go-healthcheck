# Go Health check

## Prerequisites

- go version go1.17 darwin/amd64
- Mockoon
  - require node v14.17.6
- docker [optional]

### Setup

#### Mockoon

- macOS: Use Mockoon via apps install -> `brew install --cask mockoon`

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
