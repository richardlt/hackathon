package master

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/richardlt/hackathon/types"
	"github.com/richardlt/hackathon/types/questions"
)

var clients map[string]*types.Client
var qs questions.Questions

// Serve master
func Serve() {
	clients = make(map[string]*types.Client)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Post("/register", func(c echo.Context) error {
		var register types.Register
		c.Bind(&register)
		if register.URL == "" {
			return c.JSON(http.StatusBadRequest, nil)
		}
		client := types.Client{URL: register.URL}
		// TODO check if slave already exist, prefer map
		if _, ok := clients[client.URL]; !ok {
			clients[client.URL] = &client
		}

		return c.JSON(http.StatusOK, nil)
	})
	go AskQuestions()
	e.Run(standard.New(":8080"))
}
