package v1

import (
	"net/http"
	"github.com/zenazn/goji/web"
)

func CheckSession(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}