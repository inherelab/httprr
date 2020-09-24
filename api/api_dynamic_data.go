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

// DynamicData dynamic data examples
type DynamicData struct {}

func Base64ValueGet(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

func BytesNGet(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

func DelayDelayDelete(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

func DelayDelayGet(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

func DelayDelayPatch(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

func DelayDelayPost(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

func DelayDelayPut(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

func DripGet(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

func LinksNOffsetGet(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

func RangeNumbytesGet(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

func StreamBytesNGet(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

func StreamNGet(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}

func UuidGet(c *rux.Context) {
	c.JSON(200, app.BuildReplay(c))
}
