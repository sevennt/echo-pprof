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

	// echopprof also plays well with *gin.RouterGroup
	// group := e.Group("/debug/pprof")
	// echopprof.WrapGroup(group)

	e.Start(":8080")
}
