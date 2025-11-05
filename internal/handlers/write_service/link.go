package write_service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShortLinkRequest struct {
	LongLink string `json:"long_link"`
}

func (a *WriteServiceAPI) SaveLink(c *gin.Context) {
	request := &ShortLinkRequest{}
	err := c.ShouldBindBodyWithJSON(request)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			gin.H{
				"error":   "UnprocessableEntity",
				"details": "Field long_link is required",
			},
		)
	}

	dto, err := a.controller.CreateLink(request.LongLink)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"error":   "BadRequest",
				"details": err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusOK, dto)
}
