package routes

import (
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	api_v1_controllers "github.com/takaaki-mizuno/goji-boilerplate/app/http/controllers/api/v1"
	api_v1_middlewares "github.com/takaaki-mizuno/goji-boilerplate/app/http/middlewares/api/v1"
)

func APIV1Routes() {

	apiMux := web.New()
	goji.Handle("/api/v1/*", apiMux)
	apiMux.Get("/api/v1/status", api_v1_controllers.Status_get_handler)

	apiMux.Use(api_v1_middlewares.ResponseHeader)

}
