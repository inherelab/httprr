package main

import (
	"fmt"

	"github.com/gookit/rux"
	"github.com/gookit/rux/handlers"
	"github.com/gookit/view"
	"github.com/inherelab/httpres-web/route"
)

var debug = true

func init() {
	// view templates
	view.Initialize(func(r *view.Renderer) {
		r.ViewsDir = "views"
	})
}

// start: go run main.go
// access: http://127.0.0.1:18080
func main() {
	// open debug
	rux.Debug(debug)

	r := rux.New()

	// one file
	// r.StaticFile("/site.js", "testdata/site.js")
	// allow any files in the dir.
	r.StaticDir("/static", "static")
	// add file ext limit
	// r.StaticFiles("", "testdata", "css|js")
	// r.StaticFiles("/static", "testdata", "css|js|json")

	// fmt.Println(r)
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
