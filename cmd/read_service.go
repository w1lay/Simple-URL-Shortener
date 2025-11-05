package main

import (
	"net"

	"github.com/sirupsen/logrus"
	"github.com/w1lay/Simple-URL-Shortener/internal/handlers/read_service"
	"github.com/w1lay/Simple-URL-Shortener/internal/handlers/server"
	"github.com/w1lay/Simple-URL-Shortener/internal/settings"
	"github.com/w1lay/Simple-URL-Shortener/internal/utils"
)

func main() {
	utils.SetupSettingsAndLogger()

	api := read_service.NewReadServiceAPI()
	router := api.InitRouter() // Роутер GIN выступает в качестве http.Handler
	apiUrl := net.JoinHostPort(
		settings.Settings.API.Host,
		settings.Settings.API.Port,
	)
	logrus.Infof("Running API at %s, [%s mode]", apiUrl, settings.Settings.App.Mode)

	srv := server.NewServer(router)
	err := srv.ListenAndServe()
	if err != nil {
		logrus.Fatal(err)
	}
}
