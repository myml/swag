package ui

import (
	"embed"
	"net/http"
	"strings"

	"github.com/myml/swag/swagger"
)

//go:embed index.html
//go:embed swagger-ui-bundle_3.52.5.js
//go:embed swagger-ui_3.52.5.css
var static embed.FS

func Handler(prefix string, api *swagger.API) http.HandlerFunc {
	server := http.StripPrefix(prefix, http.FileServer(http.FS(static)))
	return func(rw http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "swagger.json") {
			api.Handler(false)(rw, r)
			return
		}
		server.ServeHTTP(rw, r)
	}
}
