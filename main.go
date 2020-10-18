package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/gookit/gcli/v2"
	"github.com/gookit/rux"
	"github.com/gookit/rux/handlers"
	"github.com/gookit/view"
	"github.com/inherelab/httprr/app"
	"github.com/inherelab/httprr/route"
)

var (
	port uint
	debug = os.Getenv("HTTPRR_DEBUG") == "true"
)

func bindFlags()  {
	binName := os.Args[0]
	defDebug := os.Getenv("HTTPRR_DEBUG") == "true"

	gf := gcli.NewFlags()
	gf.FSet().Usage = func() {
		color.Infoln("Usage:", binName, "[--OPTIONS] ...\n")
		color.Yellow.Println("Options:")
		gf.PrintHelpPanel()
	}
	gf.BoolOpt(&debug, "debug", "", defDebug, "open debug mode")
	gf.UintOpt(&port, "port", "p", 18081, "the http server port")

	err := gf.Parse(os.Args[1:])
	if err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}

		panic(err)
	}

	app.HttpAddr = fmt.Sprintf("127.0.0.1:%d", port)
}

func init() {
	bindFlags()

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
	r.Listen(app.HttpAddr)
	// apply global pre-handlers
	// http.ListenAndServe(":18080", handlers.HTTPMethodOverrideHandler(r))
}
