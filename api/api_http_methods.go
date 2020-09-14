/*
 * httprr.trys.cc
 *
 * A simple HTTP Request & Response Service.<br/> <br/> <b>Run locally: </b> <code>$ docker run -p 80:80 inherelab/httprr</code>
 *
 * API version: 0.0.1
 * Contact: in.798@qq.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"github.com/gookit/rux"
	"github.com/inherelab/httprr/app"
)

// HttpMethods APIs http methods examples
type HttpMethods struct {}

func (grp *HttpMethods) AddRoutes(r *rux.Router) {
	r.GET("get", grp.Get)
	r.PUT("put", grp.Put)
	r.POST("post", grp.Post)
	r.PATCH("patch", grp.Patch)
	r.DELETE("delete", grp.Delete)
}

func (*HttpMethods) Delete(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

func (*HttpMethods) Get(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

func (*HttpMethods) Patch(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

func (*HttpMethods) Post(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

func (*HttpMethods) Put(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}
