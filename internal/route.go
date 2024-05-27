package internal

import (
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, handler *Handler) {
	e.POST("/clients", handler.createClient)
	e.GET("/clients/:id", handler.getClient)
	e.GET("/clients", handler.getClients)
	e.GET("/clients/assign-lead", handler.assignLead)
}
