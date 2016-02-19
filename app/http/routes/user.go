package routes

import (
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	user_controller "github.com/takaaki-mizuno/goji-boilerplate/app/http/controllers/user"
)

func UserRoutes() {

	userMux := web.New()
	goji.Handle("/*", userMux)
	userMux.Get("/", user_controller.Index_get_handler)

}
