package main

import (
	"net"

	"github.com/sirupsen/logrus"
	controllers "github.com/w1lay/Simple-URL-Shortener/internal/controllers/write_service"
	"github.com/w1lay/Simple-URL-Shortener/internal/handlers/server"
	"github.com/w1lay/Simple-URL-Shortener/internal/handlers/write_service"
	"github.com/w1lay/Simple-URL-Shortener/internal/repositories"
	"github.com/w1lay/Simple-URL-Shortener/internal/settings"
	"github.com/w1lay/Simple-URL-Shortener/internal/utils"
)

func main() {
	utils.SetupSettingsAndLogger()

	db := repositories.NewPostgresqlConnection()
	repository := repositories.NewLinkPostgresRepository(db)
	controller := controllers.NewWriteServiceController(repository)
	api := write_service.NewWriteServiceAPI(controller)

	router := api.InitRouter() // Роутер GIN выступает в качестве http.Handler
	apiUrl := net.JoinHostPort(
		settings.Settings.API.Host,
		settings.Settings.API.Port,
	)
	logrus.Infof("Running WriteService API at %s, [%s mode]", apiUrl, settings.Settings.App.Mode)

	srv := server.NewServer(router)
	err := srv.ListenAndServe()
	if err != nil {
		logrus.Fatal(err)
	}
}
