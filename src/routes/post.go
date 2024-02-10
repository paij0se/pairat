package routes

import (
	routes "github.com/drpaij0se/pairat/src/controllers"
	"github.com/labstack/echo"
)

func Post(e *echo.Echo) {
	e.POST("/commands", routes.UploadCommand)
	e.POST("/commands/ansitrue", routes.AnsiOn)
}
