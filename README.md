# scanii-go

[![GoDoc](https://godoc.org/github.com/uvasoftware/scanii-go?status.svg)](https://godoc.org/github.com/uvasoftware/scanii-go)

scanii-go is a library for use with the [UVA Soft](http://www.uvasoftware.com/) [scanii.com](http://www.scanii.com) API.

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

This is a convenience for the actual command to run Go unit testing, `go test -v -cover -coverprofile coverate_report/cover.out ./...` .  It's best to run the `make` command.

To get a detailed report on what has bits of code have actual coverage, run the command below.  It will bring you into a browser session with explicit detail on what has coverage and what doesn't.

```Go
go tool cover -html=cover.out
```

## Examples

Build a client

```Go
clientOpts := &scaniigo.ClientOpts{
	Version:  "2.1",
	Validate: false,
}
c, err := scaniigo.NewClient(clientOpts)
if err != nil {
	log.Fatalln(err)
}
```

Verify connectivity

```Go
res, err := c.Ping()
if err != nil {
	log.Fatalln(err)
}
```

Retrieve auth token

```Go
res, err := c.RetrieveAuthToken("auth-id-string")
if err != nil {
	log.Fatalln(err)
}
```

Create a temporary auth token

```Go
res, err := c.CreateTempAuthToken()
if err != nil {
	log.Fatalln(err)
}
```

Delete a temporary auth token

```Go
if err := c.DeleteTempAuthToken("token-id"); err != nil {
	log.Fatalln(err)
}
```

Retrieve a previously processed file

```Go
res, err := c.RetrieveProcessedFile("file-id-string")
if err != nil {
	log.Fatalln(err)
}
```

Process local file synchronously

```Go
pfp := &scaniigo.ProcessFileParams{
	File: "fib",
}
pfr, err := c.ProcessFileSync(pfp)
if err != nil {
	log.Fatalln(err)
}
```

Process local file asynchronously

```Go
pfp := &scaniigo.ProcessFileParams{
	File: "fib",
}
pfr, err := c.ProcessFileAsync(pfp)
if err != nil {
	log.Fatalln(err)
}
```

Process a remote file asynchronously

```Go
res, err := c.ProcessRemoteFileAsync(rfap)
if err != nil {
	log.Fatall(err)
}
```
