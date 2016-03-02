# scanii-go

WIP 

[![GoDoc](https://godoc.org/github.com/uvasoftware/scanii-go?status.svg)](https://godoc.org/github.com/uvasoftware/scanii-go)

scanii-go is a library for use with the [UVA Soft](http://www.uvasoftware.com/) [scanii.com](http://www.scanii.com) API.

Examples can be found in the `_examples` directory.

## Documentation

Details about this package can be viewed by clicking on the "GoDoc" badge above, by running the local Go doc server, or by referencing the code comments.

Run the local doc server.
```sh
$ godoc -http=":6060"
```

## Install

```sh
$ go get github.com/uvasoftware/scanii-go
```

## Development

When the API is expanded or altered, one can simply add or adjust the correlated method mapping.

Field validation on parameter types is done by passing the given type to a Validate() function.  This function expects a parameter of type Validator which is an interface that implements a function with the signature of `Validate() error`.  

### Dependencies

```sh
$ make dep
```

### Testing

Test all changes by running the command below.  

```sh
make test
```

This is a convenience for the actual command to run Go unit testing, `go test -v -cover ./...` .  It's best to run the `make` command.
