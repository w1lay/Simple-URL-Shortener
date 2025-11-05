package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/w1lay/Simple-URL-Shortener/internal/settings"
)

func SetupSettingsAndLogger() {
	settings.Init()

	switch settings.Settings.App.Mode {
	case "dev":
		logrus.SetFormatter(&logrus.TextFormatter{})
		gin.SetMode(gin.DebugMode)
	case "prod":
		logrus.SetFormatter(&logrus.JSONFormatter{})
		gin.SetMode(gin.ReleaseMode)
	default:
		logrus.SetFormatter(&logrus.JSONFormatter{})
		gin.SetMode(gin.ReleaseMode)
	}

	gin.DebugPrintFunc = logrus.Infof
	logrus.Infof("Settings and logger was initialized")
}
