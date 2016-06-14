package slave

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

// Serve slave
func Serve() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Get("/answer", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})

	e.Run(standard.New(":8081"))
}
