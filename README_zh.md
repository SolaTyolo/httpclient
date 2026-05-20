# httpclient

[English](README.md)

基于 [go-retryablehttp](https://github.com/hashicorp/go-retryablehttp) 的可重试 JSON HTTP 客户端。**所有请求均经 `*retryablehttp.Client` 发出**，不提供单独的 `net/http` 传输路径。

## 安装

```bash
go get github.com/SolaTyolo/httpclient
```

需要 Go 1.19+。

## 目录结构

```
httpclient/
├── client.go      # HTTP、New、Requester
├── option.go      # ClientOption、With*
├── request.go     # Get / Post / Put / Delete
├── transport/     # Requester 接口
├── retry/         # 默认参数、重试策略、NewClient
├── auth/          # Bearer / Basic
├── response/      # Response、StatusError、DecodeJSON
└── endpoint/      # URL 与查询参数
```

日常使用只需 `import "github.com/SolaTyolo/httpclient"`；仅需某一能力时可引用对应子包。

## 快速开始

```go
import (
    "github.com/SolaTyolo/httpclient"
    "github.com/SolaTyolo/httpclient/response"
)

h := httpclient.New(httpclient.WithTimeout(30 * time.Second))
resp, err := h.Post(ctx, url, body)
if err != nil {
    var se *response.StatusError
    if errors.As(err, &se) { /* se.StatusCode、se.Body */ }
    return err
}
var out MyType
return resp.DecodeJSON(&out)
```

默认：**3 次重试**、**5s** 超时、退避 **2–10s**、**5xx** 重试。

## 选项

| 选项 | 说明 |
|------|------|
| `WithTimeout` | 单次请求超时 |
| `WithRetryMax` | 最大重试次数 |
| `WithRetryWait` | 退避间隔 |
| `WithCheckRetry` | `retry.DefaultRetryFn`、`retry.NoCheckRetryFn` 等 |
| `WithAuthFn` | 请求前钩子 |
| `WithRetryableClient` | 注入 `*retryablehttp.Client` |
| `WithRequester` | 同上（接口形式） |

## 自定义 retryablehttp

```go
rc := retryablehttp.NewClient()
rc.HTTPClient = &http.Client{Transport: myTransport}
h := httpclient.New(httpclient.WithRetryableClient(rc))
```

## 认证

```go
import "github.com/SolaTyolo/httpclient/auth"

a := auth.New("token")
h := httpclient.New(httpclient.WithAuthFn(a.BearAuth()))
```

## 开发

```bash
make test
```

## 许可证

MIT — 见 [LICENSE](LICENSE)。
