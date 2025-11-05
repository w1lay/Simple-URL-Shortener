package write_service

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/w1lay/Simple-URL-Shortener/internal/controllers/write_service"
	"github.com/w1lay/Simple-URL-Shortener/internal/handlers/middlewares"
)

type WriteServiceAPI struct {
	controller *controllers.WriteServiceController
}

func NewWriteServiceAPI(controller *controllers.WriteServiceController) *WriteServiceAPI {
	return &WriteServiceAPI{controller: controller}
}

func (a *WriteServiceAPI) InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())

	router.GET("/", a.Healtcheck)
	router.POST("/link/", a.SaveLink)
	return router
}
