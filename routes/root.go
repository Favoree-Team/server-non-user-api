package routes

import (
	"github.com/Favoree-Team/server-non-user-api/config"
)

var (
	DB = config.ConnectDB()
)
