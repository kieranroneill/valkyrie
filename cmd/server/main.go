package main

import (
	"github.com/kieranroneill/valkyrie/pkg/application"
	"github.com/kieranroneill/valkyrie/pkg/cleanup"
  "github.com/kieranroneill/valkyrie/pkg/logger"
	"github.com/kieranroneill/valkyrie/pkg/router"
	"github.com/kieranroneill/valkyrie/pkg/server"
)

func main() {
	app, err := application.New()
	if err != nil {
		logger.Error.Fatal(err.Error())
	}

	srv := server.
    New().
		WithAddr(":" + app.Config.Port).
		WithRouter(router.New(app)).
		WithErrLogger(logger.Error)

	go func() {
		logger.Info.Printf("🚀 blast off in %s on :%s", app.Config.Environment, app.Config.Port)

		if err := srv.Start(); err != nil {
			logger.Error.Fatal(err.Error())
		}
	}()

	cleanup.Init(func() {
		if err := srv.Close(); err != nil {
			logger.Error.Println(err.Error())
		}
	})
}
