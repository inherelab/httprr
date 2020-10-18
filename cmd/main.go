package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
	"github.com/gookit/gcli/v2"
	"github.com/gookit/slog"
)

var (
	swagFile string
	outFile string
	apiHost string
)

var upSwagDoc = &gcli.Command{
	Name:   "updoc",
	UseFor: "replace real host to the swagger doc file",
	Config: func(c *gcli.Command) {
		c.StrOpt(&swagFile,
			"swagger-file", "f",
			"resource/swagger.json",
			"the swagger document file path",
		)
		c.StrOpt(&outFile,
			"outfile", "o",
			"static/swagger.json",
			"the new swagger document file output path",
		)
		c.StrOpt(&apiHost,
			"host", "H",
			"127.0.0.1:18081",
			"the swagger document server host",
		)
	},
	Examples: "{$fullCmd} -H hrr.trys.cc",
	Func: func(c *gcli.Command, _ []string) error {
		slog.Println("read file: ", swagFile)
		bts, err := ioutil.ReadFile(swagFile)
		if err != nil {
			return err
		}

		text := strings.Replace(string(bts), "{{host}}", apiHost, -1)

		outFile := "static/" + filepath.Base(swagFile)
		slog.Println("write file: ", outFile)
		err = ioutil.WriteFile(outFile, []byte(text), 0664)
		if err == nil {
			color.Info.Println("operate successful")
		}

		return err
	},
}

func main() {
	app := gcli.NewApp()

	app.AddCommand(upSwagDoc)

	app.Run()
}
