package route

import (
	"github.com/gookit/rux"
	"github.com/inherelab/httpres-web/web"
)

func LoadRoutes(r *rux.Router) {
	r.Controller("/", &web.HomeController{})

	// apis
}
