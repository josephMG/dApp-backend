package bootstrap

import (
	"hardhat-backend/api/controllers"
	"hardhat-backend/api/middlewares"
	"hardhat-backend/api/routes"
	"hardhat-backend/config"
	"hardhat-backend/infrastructure"
	"hardhat-backend/lib"
	"hardhat-backend/lib/loggers"
	"hardhat-backend/repository"
	"hardhat-backend/services"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
	config.Module,
	loggers.Module,
	services.Module,
	infrastructure.Module,
	middlewares.Module,
	lib.Module,
	repository.Module,
)
