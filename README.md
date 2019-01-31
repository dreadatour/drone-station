# Drone Station

Drone station used to track Drones.


## Drone coordinate system

Drone coordinates defined with absolute `quadrant` value and pair of `(x, y)` coordinates relative to quadrant.

To address quadrants [Geoshash](https://en.wikipedia.org/wiki/Geohash) is used.

X and Y coordinates are relative to quadrant, so X = `0.0` means quadrant minimal longitude (west boundary) and X = `100.0` means quadrant maximal longitude (east boundary). Same for Y (Y = `0.0` for south boundary, Y = `100.0` for north boundary). Values less than 0 and greater than 100 will be rejected.

For example, coordinate

```
latitude=51.924333714
longitude=4.477883541
```

will be treated as

```
quadrant=u15pmus9
x=78.11
y=16.75
```


## Config

All configuration taken from environment variables.

Available variables and their values:

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

or simply:

```
$ make api
```

### Documentation

API documentation leaves in `docs` folder in OpenAPI format. To view API documentation, simply run:

```
$ make open_api_docs
```

### Vendoring

Third-party packages are vendored using [dep](https://github.com/golang/dep).
