package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewUserRoutes),
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewContractRoutes),
	fx.Provide(NewSiweRoutes),
	fx.Provide(NewRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	userRoutes *UserRoutes,
	authRoutes *AuthRoutes,
	contractRoutes *ContractRoutes,
	siweRoutes *SiweRoutes,
) Routes {
	return Routes{
		userRoutes,
		authRoutes,
		contractRoutes,
		siweRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
