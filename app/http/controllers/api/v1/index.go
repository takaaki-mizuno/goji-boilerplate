package v1

import (
	"net/http"
	"github.com/zenazn/goji/web"

	services "github.com/takaaki-mizuno/goji-boilerplate/app/services"
	api_v1_response "github.com/takaaki-mizuno/goji-boilerplate/app/http/responses/api/v1"

)

func Status_get_handler(c web.C, w http.ResponseWriter, r *http.Request) {
	services.Api().GenerateResponse(c, r, w, api_v1_response.Status{Status:"ok"})
}
