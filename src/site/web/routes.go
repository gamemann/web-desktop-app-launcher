package web

import (
	"net/http"

	"github.com/gamemann/web-desktop-app-launcher/config"
)

func SetupRoutes(cfg *config.Config) {
	// Serve static mages.
	fs := http.FileServer(http.Dir("images"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))

	// Serve JavaScript.
	fs = http.FileServer(http.Dir("js"))
	http.Handle("/js/", http.StripPrefix("/js/", fs))

	// Serve CSS.
	fs = http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	// Setup root handler.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RootHandler(w, r, cfg)
	})

	// Setup back-end handler.
	http.HandleFunc("/backend/submit", func(w http.ResponseWriter, r *http.Request) {
		BackendHandler(w, r, cfg)
	})
}
