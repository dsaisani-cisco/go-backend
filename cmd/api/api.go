package api

import (
	"net/http"
	"time"

	"danielsaisani/cmd/service"

	"github.com/gin-gonic/gin"
)

type API struct {
	timeService *service.TimeService
}

func New(service *service.TimeService) *API {
	return &API{
		timeService: service,
	}
}

// RegisterRoutes mounts the API's endpoints onto the given router.
func (a *API) RegisterRoutes(router gin.IRouter) {
	router.GET("/now", a.now)
}

type nowResponse struct {
	Now time.Time `json:"now"`
}

// now returns..... the current time, what else did you expect?
func (a *API) now(c *gin.Context) {
	c.JSON(http.StatusOK, nowResponse{Now: a.timeService.Now().UTC()})
}
