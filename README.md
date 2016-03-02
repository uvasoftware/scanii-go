# scanii-go

[![GoDoc](https://godoc.org/github.com/uvasoftware/scanii-go?status.svg)](https://godoc.org/github.com/uvasoftware/scanii-go)

scanii-go is a library for use with the [UVA Soft](http://www.uvasoftware.com/) [scanii.com](http://www.scanii.com) API.

Examples can be found in the `_examples` directory.

## Documentation

Details about this package can be viewed by clicking on the "GoDoc" badge above, by running the local Go doc server, or by referencing the code comments.

Run the local doc server.
```sh
$ godoc -http=":6060"
```

## Features

## Install

```sh
$ go get github.com/uvasoftware/scanii-go
```

## Development

### Dependencies

```sh
$ make dep
```

### Testing

Test all changes by running the command below.  This is a convenience for the actual command to run Go unit testing, `go test -v -cover ./...`

```sh
make test
```
