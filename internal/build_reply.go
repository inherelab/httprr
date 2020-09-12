package internal

import (
	"net/http"
	"strings"

	"github.com/gookit/rux"
	"github.com/inherelab/httpres-web/app"
)

func BuildBodyReplay(c *rux.Context) app.BodyModel {


	bm := app.BodyModel{
		Url:     c.URL().String(),
		Origin:  "",
		Body:    c.Req.Body,
		Json:    nil,
		Args:    nil,
		Form:    nil,
		Files:   nil,
		Headers: convHeaders2map(c.Req.Header),
	}

	return bm
}

func BuildQueryReplay(c *rux.Context) app.QueryModel {
	qm := app.QueryModel{
		Url:     c.URL().String(),
		Args:    nil,
		Origin:  "",
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