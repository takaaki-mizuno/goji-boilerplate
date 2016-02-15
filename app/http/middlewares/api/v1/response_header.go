package v1

import (
	"net/http"
	"github.com/zenazn/goji/web"
)

func ResponseHeader(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		w.Header().Set("X-BoilerPlate-Version", "1.0")
	}

	return http.HandlerFunc(fn)
}
