package routes

import (
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	api_v1_controllers "github.com/takaaki-mizuno/goji-boilerplate/app/http/controllers/api/v1"
	api_v1_middlewares "github.com/takaaki-mizuno/goji-boilerplate/app/http/middlewares/api/v1"
)

func APIV1Routes(mainMux *web.Mux) {
	goji.Handle("/api/v1/*", mainMux)
	mainMux.Get("/api/v1/status", api_v1_controllers.Status_get_handler)

	mainMux.Use(api_v1_middlewares.ResponseHeader)
}
