# go-server

Generic HTTP server utilities for Go applications using Gin.

## Features

* **Gin Engine**: Creates and configures a Gin engine
* **Middleware**: Composable middleware through functional options
* **Trusted Proxies**: Configure trusted proxy addresses
* **Server Lifecycle**: HTTP server startup and graceful shutdown

## Installation

```bash
go get github.com/kingzslayer/go-server
```

## Create an Engine

```go
import "github.com/kingzslayer/go-server"

engine, err := server.NewEngine(
    server.WithMiddleware(
        recoveryMiddleware,
        loggingMiddleware,
        corsMiddleware,
    ),
    server.WithTrustedProxies([]string{"127.0.0.1"}),
)

if err != nil {
    log.Fatal(err)
}
```

## Create a Server

`go-server` accepts a standard `http.Handler`, so the Gin engine can be passed directly.

```go
cfg := server.ServerConfig{
    Addr:            ":8080",
    ReadTimeout:     15 * time.Second,
    WriteTimeout:    15 * time.Second,
    IdleTimeout:     60 * time.Second,
    ShutdownTimeout: 10 * time.Second,
}

srv := server.NewServer(cfg, engine)
```

## Run

```go
if err := srv.Run(context.Background()); err != nil {
    log.Fatal(err)
}
```

`Run`:

* Starts the HTTP server
* Listens for `SIGINT` and `SIGTERM`
* Gracefully shuts down the server
* Uses the configured shutdown timeout

## Complete Example

```go
package main

import (
    "context"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    server "github.com/kingzslayer/go-server"
)

func main() {
    engine, err := server.NewEngine(
        server.WithMiddleware(
            func(c *gin.Context) {
                c.Next()
            },
        ),
    )
    if err != nil {
        log.Fatal(err)
    }

    engine.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
        })
    })

    cfg := server.ServerConfig{
        Addr:            ":8080",
        ReadTimeout:     15 * time.Second,
        WriteTimeout:    15 * time.Second,
        IdleTimeout:     60 * time.Second,
        ShutdownTimeout: 10 * time.Second,
    }

    srv := server.NewServer(cfg, engine)

    if err := srv.Run(context.Background()); err != nil {
        log.Fatal(err)
    }
}
```
