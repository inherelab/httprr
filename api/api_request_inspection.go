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

// RequestInspection Inspect the request data
type RequestInspection struct {}

func (grp *RequestInspection) AddRoutes(r *rux.Router) {
	r.GET("headers", grp.Headers)
	r.GET("ip", grp.Ip)
	r.GET("user-agent", grp.UserAgent)
}

func (*RequestInspection) Headers(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

func (*RequestInspection) Ip(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

func (*RequestInspection) UserAgent(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}
