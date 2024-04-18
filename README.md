Httpclient
===========

HTTP Client Based on go-retryablehttp


Example Use
===========

Using this library should look almost identical to what you would do with
`net/http`. The most simple example of a GET request is shown below:

```go
resp, err := httpclient.Default().Get("/foo")
if err != nil {
    panic(err)
}
```


For more usage and examples see the
[godoc](https://pkg.go.dev/github.com/SolaTyolo/httpclient).