package route

import (
	"github.com/gookit/rux"
	"github.com/inherelab/httprr/api"
	"github.com/inherelab/httprr/web"
)

func LoadRoutes(r *rux.Router) {
	r.Controller("/", &web.HomeController{})

	// apis
	r.Controller("/", &api.HttpMethods{})

	// r.Controller("/", &api.RequestInspection{})
	ri := &api.RequestInspection{}
	r.AddRoute(rux.NewRoute("/headers", ri.Headers, rux.GET))
	r.AddRoute(rux.NewRoute("/user-agent", ri.UserAgent, rux.GET))
	r.AddRoute(rux.NewRoute("/ip", ri.Ip, rux.GET))

	r.Controller("/status", &api.StatusCode{})
	r.Controller("/anything", &api.Anything{})
}
