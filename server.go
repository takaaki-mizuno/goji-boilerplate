package main

import (
	"github.com/takaaki-mizuno/goji-boilerplate/app"

	"github.com/zenazn/goji"
	"github.com/hypebeast/gojistatic"
)

func main() {
	goji.Use(gojistatic.Static("public", gojistatic.StaticOptions{SkipLogging: true}))
	app.App()
	goji.Serve()
}
