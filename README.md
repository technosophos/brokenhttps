# Broken HTTPS

This starts an HTTP server with an invalid certificate. It is useful for testing how your app handles HTTPS (TLS/SSL) when the encryption is fine but the certificate should not be trusted.

## Usage

Use `go get` or `git clone` to get this repo.

To run: `go run server.go`. If you would like to use your own data payload, you can supply a full path to the file it should read: `go run server.go /path/to/my/file.txt`.

To stop, use `CTRL-C`.