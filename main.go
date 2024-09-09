package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/nilssonr/agentside-login/web/page"
)

//go:embed static
var staticFS embed.FS

func main() {
	r := chi.NewRouter()

	r.Handle("/static/*", http.FileServer(http.FS(staticFS)))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(page.Login()).ServeHTTP(w, r)
	})

	r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(page.Logout()).ServeHTTP(w, r)
	})

	r.Get("/consent", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(page.Consent()).ServeHTTP(w, r)
	})

	if err := http.ListenAndServe(":4010", r); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
