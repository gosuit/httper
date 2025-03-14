# Httper

This is a simple and flexible HTTP library for Go that provides an easy way to create HTTP servers and clients. The library allows you to configure server settings, manage requests, and handle responses with support for marshaling and unmarshaling of request and response bodies.

## Installation

```zsh
go get github.com/gosuit/httper
```

## Features

- Create and configure HTTP servers with customizable timeouts.
- Build HTTP clients with URL prefixes and timeout settings.
- Easy-to-use API for making HTTP requests.
- Support for marshaling and unmarshaling request and response bodies in JSON, XML and other formats.

## Usage

### Server

```golang
package main

import (
    "time"
    "log/slog"

    "github.com/gosuit/httper"
)


func main() {
    log := slog.Default()

    cfg := httper.ServerCfg{
        Url:             ":8080",
        ReadTimeout:     5 * time.Second,
        WriteTimeout:    10 * time.Second,
        ShutdownTimeout: 5 * time.Second,
    }

    handler := http.HandlerFunc(yourHandlerFunction) // Your custom handler function
    server := httper.NewServer(cfg, handler)

    server.Start()

	err := server.Shutdown(log)
    // Handling error
}
```

### Client

```golang
package main

import (
    "time"
    "log/slog"

    "github.com/gosuit/httper"
)

type Result struct {
    Field1 string `json:"field"`
}

func main() {
    log := slog.Default()

    clientCfg := httper.ClientCfg{
        Prefix:  "http://api.example.com",
        Timeout: 5 * time.Second,
    }

    client := httper.NewClient(clientCfg)  
    
    var result Result

    req, err := httper.NewReq(&httper.Params{
        Method:        httper.GetMethod,
        Url:           "/endpoint",
        Unmarshal:     true,
        UnmarshalTo:   &result,
        UnmarshalType: httper.JsonType,
    })
    if err != nil {
        log.Fatal(err)
    }

    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }

    // Handle response
    fmt.Println(result)
}
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.