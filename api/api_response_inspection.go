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

// CacheGet operate API
func CacheGet(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

// CacheValueGet operate API
func CacheValueGet(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

// EtagEtagGet operate API
func EtagEtagGet(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

// ResponseHeadersGet operate API
func ResponseHeadersGet(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

// ResponseHeadersPost operate API
func ResponseHeadersPost(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}
