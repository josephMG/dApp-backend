package bootstrap

import (
	"hardhat-backend/api/controllers"
	"hardhat-backend/api/middlewares"
	"hardhat-backend/api/routes"
	"hardhat-backend/config"
	"hardhat-backend/lib"
	"hardhat-backend/lib/loggers"
	"hardhat-backend/repository"
	"hardhat-backend/services"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	config.Module,
	loggers.Module,
	services.Module,
	middlewares.Module,
	repository.Module,
)
