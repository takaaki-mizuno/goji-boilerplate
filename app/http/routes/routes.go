package routes

import (
	"github.com/zenazn/goji/web"
)

func Routes() {
	mainMux := web.New()
	AdminRoutes(mainMux)
	UserRoutes(mainMux)
	APIV1Routes(mainMux)
}

