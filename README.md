# Drone Station

Drone station used to track Drones.


## Config

All configuration taken from environment variables. Available variables and their values.

### HTTP server configuration

* *HTTP_HOST* is HTTP server listen IP address (examples: `127.0.0.1` for local interface only, `0.0.0.0` for all interfaces)
* *HTTP_PORT* is HTTP server port (examples: `80`, `8080`)

### Logging configuration

All logs goes into STDOUT. Configuration:

* *LOG_LEVEL* is logging level (available values are pretty common: `debug`, `info`, `warning`, `error`, `fatal` and `panic`).
* *LOG_FORMAT* is log output format: `json` for JSON logs, `text` for colorised text logs and `plaintext` for plain text logs.


## Development

### Config

You can use configuration from `.env` file during development. To do that use `-dotenv` flag while running server:

```
$ go run cmd/api/main.go -dotenv
```

### Vendoring

Third-party packages are vendored using [dep](https://github.com/golang/dep).
