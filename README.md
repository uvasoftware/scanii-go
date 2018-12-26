### A pure Go interface to the Scanii content processing service - https://scanii.com

### How to use this client

#### Installing

```sh
$ go get github.com/uvasoftware/scanii-go 
```

### Sample usage:
 
```go
client := scanii.NewClient(&scanii.ClientOpts{
		Target: endpoints.LATEST,
		Key:    key,
		Secret: secret,
	})

file, _ := ioutil.TempFile("", "")
defer func() {
    _ = file.Close()
    _ = os.Remove(file.Name())
}()

_, _ = file.Write([]byte("hello world"))

metadata := map[string]string{
    "m1": "v1",
    "m2": "v2",
}

r, err := client.Process(file.Name(), "", metadata)

```

Please note that you will need a valid scanii.com account and API Credentials. 

More advanced usage examples can be found [here](https://github.com/uvasoftware/scanii-go/blob/master/integration_test.go)

General documentation on scanii can be found [here](http://docs.scanii.com)
 

