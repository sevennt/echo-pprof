echo-pprof
========

A wrapper for [golang web framework echo](https://github.com/labstack/echo) to use `net/http/pprof` easily.

## Install

First install echo-pprof to your GOPATH using `go get`:

```sh
go get github.com/sevenNt/echo-pprof
```

## Usage

```go
package main

import (
	"github.com/labstack/echo"
	"github.com/sevenNt/echo-pprof"
)

func main() {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	// automatically add routers for net/http/pprof
	// e.g. /debug/pprof, /debug/pprof/heap, etc.
	echopprof.Wrap(e)

	// echopprof also plays well with *echo.Group
	// prefix := "/debug/pprof"
	// group := e.Group(prefix)
	// echopprof.WrapGroup(prefix, group)

	e.Start(":8080")
}
```

Start this server, and then visit [http://127.0.0.1:8080/debug/pprof/](http://127.0.0.1:8080/debug/pprof/) and you'll see what you want.

Have Fun.
