package slave

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/richardlt/hackathon/types"
)

// Serve slave
func Serve(master string, slavePort string) {

	local := "http://localhost:" + slavePort
	fmt.Println("Local is " + local)
	fmt.Println("Master is " + master)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Post("/answer", func(c echo.Context) error {
		var question types.Question
		c.Bind(&question)

		answer, err := Answer(question)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.Answer{Value: ":/"})
		}
		return c.JSON(http.StatusOK, answer)
	})

	u := types.Register{URL: local}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	res, err := http.Post(master+"/register", "application/json; charset=utf-8", b)
	if err != nil {
		fmt.Println("Can't register to master : " + master)
		panic(err)
	} else {
		fmt.Println("Sent register to master : " + master)
	}
	defer res.Body.Close()
	e.Run(standard.New(":" + slavePort))
}
