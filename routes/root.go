package routes

import (
	"github.com/Favoree-Team/server-non-user-api/auth"
	"github.com/Favoree-Team/server-non-user-api/config"
	"github.com/Favoree-Team/server-non-user-api/middleware"
)

var (
	DB             = config.ConnectDB()
	authService    = auth.NewAuthService()
	MainMiddleware = middleware.Middleware(authService)
)
