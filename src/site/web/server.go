package web

import (
	"fmt"
	"net/http"

	"github.com/gamemann/web-desktop-app-launcher/config"
)

func SetupServer(cfg *config.Config) error {
	var err error

	// Parse host and port as address.
	addr := fmt.Sprintf("%s:%d", cfg.Web.Host, cfg.Web.Port)

	// Setup routes and handlers.
	SetupRoutes(cfg)

	err = http.ListenAndServe(addr, nil)

	return err
}
