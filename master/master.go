package master

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/richardlt/hackathon/types"
)

var clients []types.Client

// Serve master
func Serve() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Post("/register", func(c echo.Context) error {
		var register types.Register
		c.Bind(&register)
		if register.Url == "" {
			return c.JSON(http.StatusBadRequest, nil)
		}
		return c.JSON(http.StatusOK, nil)
	})

	e.Run(standard.New(":8080"))
}
