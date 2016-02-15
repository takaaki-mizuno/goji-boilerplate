package routes

import (
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	admin_controller "github.com/takaaki-mizuno/goji-boilerplate/app/http/controllers/admin"
)

func AdminRoutes(mainMux *web.Mux) {
	goji.Handle("/admin/*", mainMux)
	mainMux.Get("/admin/", admin_controller.Index_get_handler)
}
