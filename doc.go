// Package httpclient provides a retryable HTTP client for JSON APIs.
//
// Subpackages hold shared types and helpers:
//
//   - [github.com/SolaTyolo/httpclient/transport] — Requester
//   - [github.com/SolaTyolo/httpclient/retry] — defaults and retry policies
//   - [github.com/SolaTyolo/httpclient/auth] — authentication
//   - [github.com/SolaTyolo/httpclient/response] — Response, StatusError
//   - [github.com/SolaTyolo/httpclient/endpoint] — URL builder
//
// All HTTP traffic uses [github.com/hashicorp/go-retryablehttp].
//
//	h := httpclient.New(httpclient.WithTimeout(30 * time.Second))
//	resp, err := h.Post(ctx, url, body)
package httpclient
