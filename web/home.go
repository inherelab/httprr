package web

import (
	"os"

	"github.com/gookit/rux"
	"github.com/gookit/view"
	"github.com/inherelab/httpres-web/app"
)

type HomeController struct {}

func (c *HomeController) AddRoutes(r *rux.Router) {
	r.GET("", c.ApiDoc)
	r.GET("apidoc", c.ApiDoc)
	r.GET("about[.html]", c.About)
}

func (*HomeController) About(c *rux.Context) {
	c.JSON(200, rux.M{"hello": "welcome"})
}

func (*HomeController) ApiDoc(c *rux.Context) {
	fInfo, err := os.Stat("static/swagger.json")
	if err != nil {
		c.AbortWithStatus(404, "swagger doc file not exists")
		return
	}

	data := map[string]string{
		"EnvName":     "prod",
		"AppName":    "http-response",
		"JsonFile":   "/static/swagger.json",
		"SwgUIPath":  "/static/swagger-ui",
		"AssetPath":  "/static",
		"UpdateTime": fInfo.ModTime().Format(app.BaseDate),
	}

	// c.HTML(200, nil)
	c.AddError(view.Partial(c.Resp, "swagger.tpl", data))
}
