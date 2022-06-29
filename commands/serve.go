package commands

import (
	"hardhat-backend/api/middlewares"
	"hardhat-backend/api/routes"
	"hardhat-backend/config"
	"hardhat-backend/lib"
	"hardhat-backend/lib/loggers"

	"github.com/spf13/cobra"
)

// ServeCommand test command
type ServeCommand struct{}

func (s *ServeCommand) Short() string {
	return "serve application"
}

func (s *ServeCommand) Setup(cmd *cobra.Command) {}

func (s *ServeCommand) Run() lib.CommandRunner {
	return func(
		middleware middlewares.Middlewares,
		env config.Env,
		router lib.RequestHandler,
		route routes.Routes,
		logger loggers.Logger,
		database lib.Database,
	) {
		middleware.Setup()
		route.Setup()

		logger.Info("Running server")
		if env.ServerPort == "" {
			_ = router.Gin.Run()
		} else {
			_ = router.Gin.Run(":" + env.ServerPort)
		}
	}
}

func NewServeCommand() *ServeCommand {
	return &ServeCommand{}
}
