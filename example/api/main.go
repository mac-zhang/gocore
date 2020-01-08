package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sunmi-OS/gocore/api"
	"github.com/urfave/cli"
	"os"
	"sort"
)

type EchoApi struct {
}

var eApi EchoApi

type Test struct {
	Test string `json:"test"`
}

func (a *EchoApi) echoStart(c *cli.Context) error {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.POST("/", func(c echo.Context) error {

		request := api.NewRequest(c)
		response := api.NewResponse(c)

		err := request.InitRawJson()
		if err != nil {
			return response.RetError(err, 400)
		}

		msg := request.Param(`test`).GetString()

		return response.RetSuccess(msg)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
	return nil
}

func main() {

	app := cli.NewApp()
	app.Name = "IOT-seanbox"
	app.Usage = "IOT-seanbox"
	app.Email = "wenzhenxi@sunmi.com"
	app.Version = "1.0.0"
	app.Usage = "IOT-seanbox"
	app.Action = eApi.echoStart

	// 指定对于的命令
	app.Commands = []cli.Command{
		{
			Name:    "api",
			Aliases: []string{"a"},
			Usage:   "api",
			Subcommands: []cli.Command{
				{
					Name:   "start",
					Usage:  "开启API-DEMO",
					Action: eApi.echoStart,
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)

}
