package route

import (
	"github.com/gookit/rux"
	"github.com/inherelab/httprr/api"
	"github.com/inherelab/httprr/web"
)

// LoadRoutes to router
func LoadRoutes(r *rux.Router) {
	r.Controller("/", &web.HomeController{})

	// anything apis
	r.Controller("/anything", &api.Anything{})

	// auth apis
	r.GET("/bearer", api.BearerGet)
	r.GET("/basic-auth/{user}/{passwd}", api.BasicAuthUserPasswdGet)
	r.GET("/digest-auth/{qop}/{user}/{passwd}", api.DigestAuthQopUserPasswdGet)
	r.GET("/digest-auth/{qop}/{user}/{passwd}/{algorithm}", api.DigestAuthQopUserPasswdAlgorithmGet)
	r.GET("/digest-auth/{qop}/{user}/{passwd}/{algorithm}/{stale_after}", api.DigestAuthQopUserPasswdAlgorithmStaleAfterGet)

	// cookies apis
	r.Controller("/cookies", &api.Cookies{})

	// dynamic data apis
	r.GET("/base64/{value}", api.Base64ValueGet)
	r.GET("/bytes/{n}", api.BytesNGet)
	r.GET("/drip", api.DripGet)
	r.GET("/links/{n}/{offset}", api.LinksNOffsetGet)
	r.GET("/range/{numbytes}", api.RangeNumbytesGet)
	r.GET("/stream-bytes/{n}", api.StreamBytesNGet)
	r.GET("/stream/{n}", api.StreamNGet)
	r.GET("/uuid", api.UuidGet)
	r.Group("/delay", func() {
		r.GET("{delay}", api.DelayDelayGet)
		r.POST("{delay}", api.DelayDelayPost)
		r.PUT("{delay}", api.DelayDelayPut)
		r.PATCH("{delay}", api.DelayDelayPatch)
		r.DELETE("{delay}", api.DelayDelayDelete)
	})

	// http methods apis
	r.Controller("/", &api.HttpMethods{})

	// image apis
	r.Controller("/image", &api.Image{})

	// redirect apis
	r.GET("/absolute-redirect/{n}", api.AbsoluteRedirectNGet)
	r.GET("/redirect/{n}", api.RedirectNGet)
	r.GET("/relative-redirect/{n}", api.RelativeRedirectNGet)
	r.Group("/redirect-to", func() {
		r.GET("", api.RedirectToGet)
		r.POST("", api.RedirectToPost)
		r.PUT("", api.RedirectToPut)
		r.PATCH("", api.RedirectToPatch)
		r.DELETE("", api.RedirectToDelete)
	})

	// request inspection apis
	r.AddRoute(rux.NewRoute("/headers", api.Headers, rux.GET))
	r.AddRoute(rux.NewRoute("/user-agent", api.UserAgent, rux.GET))
	r.AddRoute(rux.NewRoute("/ip", api.Ip, rux.GET))

	// response formats apis
	r.GET("/brotli", api.BrotliGet)
	r.GET("/deflate", api.DeflateGet)
	r.GET("/deny", api.DenyGet)
	r.GET("/encoding/utf8", api.EncodingUtf8Get)
	r.GET("/gzip", api.GzipGet)
	r.GET("/html", api.HtmlGet)
	r.GET("/json", api.JsonGet)
	r.GET("/robots.txt", api.RobotsTxtGet)
	r.GET("/xml", api.XmlGet)

	// response inspection.go apis
	r.GET("/cache", api.CacheGet)
	r.GET("/cache/{value}", api.CacheValueGet)
	r.GET("/etag/{etag}", api.EtagEtagGet)
	r.GET("/response-headers", api.ResponseHeadersGet)
	r.POST("/response-headers", api.ResponseHeadersPost)

	// status apis
	r.Controller("/status", &api.StatusCode{})
}
