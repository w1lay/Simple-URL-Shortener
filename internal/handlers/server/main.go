package server

import (
	"net"
	"net/http"

	"github.com/w1lay/Simple-URL-Shortener/internal/settings"
)

func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:           net.JoinHostPort(settings.Settings.API.Host, settings.Settings.API.Port),
		Handler:        handler,
		MaxHeaderBytes: settings.Settings.API.MaxHeaderBytes,
		ReadTimeout:    settings.Settings.API.ReadTimeout,
		WriteTimeout:   settings.Settings.API.WriteTimeout,
	}
}
