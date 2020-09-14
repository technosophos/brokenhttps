# Broken HTTPS

This starts an HTTP server with an invalid certificate. It is useful for testing how your app handles HTTPS (TLS/SSL) when the encryption is fine but the certificate should not be trusted.

## Usage

Use `go get` or `git clone` to get this repo.

To run: `go run server.go`

To stop, use `CTRL-C`.