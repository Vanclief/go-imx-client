# go-imx-client

Unofficial Golang client for the Immutable X API

Most methods are not added yet, feel free to contribute
Sponsored by www.cryptoxolos.com

## Usage ##

### Only REST
If you don't need to use the IMX SDK, you can do:

```go
import "github.com/vanclief/go-imx-client"
```

Instantiate a new client

```go
client, _ := NewClient(Ropsten,  "", "")
```

### REST + SDK

If you need to use the IMX SDK you need to

1) Run an instance of https://github.com/Vanclief/js-imx-service

2) Create a new client

```go
import "github.com/vanclief/go-imx-client"
```

Instantiate a new client

```go
client, _ := NewClient(Ropsten,  "{wallet key}", "{alchemy key}")
```

## Development ##

Run tests:

```
make test
``` 