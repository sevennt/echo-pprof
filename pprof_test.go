package echopprof

import (
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func newServer() *echo.Echo {
	e := echo.New()
	return e
}

func checkRouters(routers []*echo.Route, t *testing.T) {
	expectedRouters := map[string]string{
		"/debug/pprof/":             "IndexHandler",
		"/debug/pprof/heap":         "HeapHandler",
		"/debug/pprof/goroutine":    "GoroutineHandler",
		"/debug/pprof/block":        "BlockHandler",
		"/debug/pprof/threadcreate": "ThreadCreateHandler",
		"/debug/pprof/cmdline":      "CmdlineHandler",
		"/debug/pprof/profile":      "ProfileHandler",
		"/debug/pprof/symbol":       "SymbolHandler",
		"/debug/pprof/trace":        "TraceHandler",
		"/debug/pprof/mutex":        "MutexHandler",
	}

	for _, router := range routers {
		if (router.Method != "GET" && router.Method != "POST") || strings.HasSuffix(router.Path, "/*") {
			continue
		}
		name, ok := expectedRouters[router.Path]
		if !ok {
			t.Errorf("missing router %s", router.Path)
		}
		if !strings.Contains(router.Name, name) {
			t.Errorf("handler for %s should contain %s, got %s", router.Path, name, router.Name)
		}
	}
}

// go test github.com/sevenNt/echo-pprof -v -run=TestWrap\$
func TestWrap(t *testing.T) {
	e := newServer()
	Wrap(e)
	checkRouters(e.Routes(), t)
}

// go test github.com/sevenNt/echo-pprof -v -run=TestWrapGroup\$
func TestWrapGroup(t *testing.T) {
	for _, prefix := range []string{"/debug"} {
		e := newServer()
		g := e.Group(prefix)
		WrapGroup(prefix, g)
		checkRouters(e.Routes(), t)
	}
}
