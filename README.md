# httpclient

[中文文档](README_zh.md)

Retryable JSON HTTP client for Go, built on [go-retryablehttp](https://github.com/hashicorp/go-retryablehttp). All requests use `*retryablehttp.Client` — there is no separate `net/http`-only transport path.

## Install

```bash
go get github.com/SolaTyolo/httpclient
```

Go 1.19+.

## Layout

```
httpclient/
├── client.go      # HTTP, New, Requester
├── option.go      # ClientOption, With*
├── request.go     # Get / Post / Put / Delete
├── transport/     # Requester interface
├── retry/         # Defaults, policies, NewClient
├── auth/          # Bearer / Basic helpers
├── response/      # Response, StatusError, DecodeJSON
└── endpoint/      # URL + query builder
```

Import the root package for typical use; import subpackages when you only need one concern.

## Quick start

```go
import (
    "github.com/SolaTyolo/httpclient"
    "github.com/SolaTyolo/httpclient/response"
)

h := httpclient.New(httpclient.WithTimeout(30 * time.Second))
resp, err := h.Post(ctx, url, body)
if err != nil {
    var se *response.StatusError
    if errors.As(err, &se) { /* se.StatusCode, se.Body */ }
    return err
}
var out MyType
return resp.DecodeJSON(&out)
```

Defaults: **3 retries**, **5s** timeout, **2–10s** backoff, retry on **5xx**.

## Options

| Option | Description |
|--------|-------------|
| `WithTimeout` | Per-attempt timeout |
| `WithRetryMax` | Max retries |
| `WithRetryWait` | Backoff min/max |
| `WithCheckRetry` | `retry.DefaultRetryFn`, `retry.NoCheckRetryFn`, … |
| `WithAuthFn` | Pre-request hook |
| `WithRetryableClient` | Inject `*retryablehttp.Client` |
| `WithRequester` | Same as above (interface) |

## Custom retryablehttp client

```go
rc := retryablehttp.NewClient()
rc.HTTPClient = &http.Client{Transport: myTransport}
h := httpclient.New(httpclient.WithRetryableClient(rc))
```

## Authentication

```go
import "github.com/SolaTyolo/httpclient/auth"

a := auth.New("token")
h := httpclient.New(httpclient.WithAuthFn(a.BearAuth()))
```

## Development

```bash
make test   # vet + race, all packages
```

## License

MIT — see [LICENSE](LICENSE).
