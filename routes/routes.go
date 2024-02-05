package routes

import "net/http"

func partial(r *http.Request) string {
	return r.URL.Query().Get("partial")
}
