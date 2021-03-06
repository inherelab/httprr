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
	"net/http"

	"github.com/gookit/rux"
)

// StatusCode APIs return submitted status code
type StatusCode struct {}

func (grp *StatusCode) AddRoutes(r *rux.Router) {
	r.GET("{code}", grp.Get)
	r.PUT("{code}", grp.Put)
	r.POST("{code}", grp.Post)
	r.PATCH("{code}", grp.Patch)
	r.DELETE("{code}", grp.Delete)
}

// @Tags StatusCode
// @Summary Status Code
// @Description get app health
// @Success 201 {string} json data
// @Failure 403 body is empty
// @Router /status/{code} [delete]
func (*StatusCode) Delete(c *rux.Context)  {
	replyStatusCode(c)
}

// @Tags appApi
// @Summary Get app config
// @Param	key		query 	string	 false		"config key string"
// @Success 201 {string} json data
// @Failure 403 body is empty
// @router /status/{code} [get]
func (*StatusCode) Get(c *rux.Context)  {
	replyStatusCode(c)
}

func (*StatusCode) Patch(c *rux.Context)  {
	replyStatusCode(c)
}

func (*StatusCode) Post(c *rux.Context)  {
	replyStatusCode(c)
}

func (*StatusCode) Put(c *rux.Context)  {
	replyStatusCode(c)
}

func replyStatusCode(c *rux.Context) {
	code := c.Params.Int("code")
	if code < 100 || code > 600 {
		c.AbortWithStatus(http.StatusBadRequest, "Invalid status code")
		return
	}

	c.SetStatus(code)
}
