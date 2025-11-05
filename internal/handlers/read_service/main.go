package read_service

import (
	"github.com/gin-gonic/gin"
	"github.com/w1lay/Simple-URL-Shortener/internal/handlers/middlewares"
)

type ReadServiceAPI struct {
}

func NewReadServiceAPI() *ReadServiceAPI {
	return &ReadServiceAPI{}
}

func (a *ReadServiceAPI) InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())

	router.GET("/", a.Healtcheck)
	return router
}
