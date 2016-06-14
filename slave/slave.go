package slave

import (
<<<<<<< HEAD
=======
	"net/http"

>>>>>>> 2fc035f... Create basic master and slave
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

<<<<<<< HEAD
// Serve run a slave
=======
// Serve slave
>>>>>>> 2fc035f... Create basic master and slave
func Serve() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
<<<<<<< HEAD
=======

	e.Get("/answer", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})

>>>>>>> 2fc035f... Create basic master and slave
	e.Run(standard.New(":8081"))
}
