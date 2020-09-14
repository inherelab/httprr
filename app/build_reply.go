package app

import (
	"net/http"
	"strings"

	"github.com/gookit/rux"
)

func BuildReplay(c *rux.Context) interface{} {
	method := c.Req.Method

	// no body, query data binding. like GET DELETE OPTION ....
	if method != "POST" && method != "PUT" && method != "PATCH" {
		return BuildBodyReplay(c)
	}

	return BuildQueryReplay(c)
}

func BuildBodyReplay(c *rux.Context) BodyModel {
	bm := BodyModel{
		Url:     c.URL().String(),
		Origin:  "",
		Method:  c.Req.Method,
		Body:    "",
		Json:    nil,
		Args:    nil,
		Form:    nil,
		Files:   nil,
		Headers: convHeaders2map(c.Req.Header),
	}

	return bm
}

func BuildQueryReplay(c *rux.Context) QueryModel {
	qm := QueryModel{
		Url:     c.URL().String(),
		Args:    nil,
		Origin:  "",
		Method:  c.Req.Method,
		Headers: convHeaders2map(c.Req.Header),
	}

	return qm
}

func convHeaders2map(h http.Header) map[string]string {
	mp := make(map[string]string)
	for name, values := range h {
		mp[name] = strings.Join(values, ";")
	}

	return mp
}