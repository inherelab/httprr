package main

import (
	"fmt"
	"os"

	"github.com/gookit/rux"
	"github.com/gookit/rux/handlers"
	"github.com/gookit/view"
	"github.com/inherelab/httprr/route"
)

var debug = os.Getenv("HTTPRR_DEBUG") == "true"

func init() {
	// view templates
	view.Initialize(func(r *view.Renderer) {
		r.ViewsDir = "resource/views"
	})

	// open debug
	rux.Debug(debug)
}

// start: go run main.go
// access: http://127.0.0.1:18080
func main() {
	r := rux.New()

	r.StaticDir("/static", "static")

	// register routes
	r.Use(handlers.PanicsHandler())
	if debug {
		r.Use(handlers.RequestLogger())
	}

	// handle error
	r.OnError = func(c *rux.Context) {
		if err := c.FirstError(); err != nil {
			fmt.Println(err)
			c.HTTPError(err.Error(), 400)
			return
		}
	}

	route.LoadRoutes(r)

	// quick start
	r.Listen("127.0.0.1:18080")
	// apply global pre-handlers
	// http.ListenAndServe(":18080", handlers.HTTPMethodOverrideHandler(r))
}
