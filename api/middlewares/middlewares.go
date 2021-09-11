package middlewares

import "net/http"

func JsonMiddleware(hf http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		hf(rw, r)
	}
}
