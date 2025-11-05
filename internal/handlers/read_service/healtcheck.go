package read_service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/w1lay/Simple-URL-Shortener/internal/settings"
)

type HealthcheckResponse struct {
	Status     string `json:"status"`
	AppMode    string `json:"app_mode"`
	ApiName    string `json:"api_name"`
	ApiVersion string `json:"api_version"`
}

func (a *ReadServiceAPI) Healtcheck(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		HealthcheckResponse{
			Status:     "ok",
			AppMode:    settings.Settings.App.Mode,
			ApiName:    "short_me_read",
			ApiVersion: settings.Settings.API.Version,
		},
	)
}
