package middleware

import (
	"errors"

	"github.com/Favoree-Team/server-non-user-api/auth"
	"github.com/Favoree-Team/server-non-user-api/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Middleware(authService auth.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || len(tokenString) == 0 {
			c.AbortWithStatusJSON(401, utils.ErrorMessages(utils.ErrorUnauthorizeUser, errors.New("error user unauthorize")))
			return
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, utils.ErrorMessages(utils.ErrorUnauthorizeUser, errors.New("error user unauthorize")))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("id", claims["id"])
			c.Set("ip_address", claims["ip_address"])
			c.Set("device_access", claims["device_access"])
			c.Set("VALID_TOKEN", tokenString)

		} else {
			c.AbortWithStatusJSON(401, utils.ErrorMessages(utils.ErrorUnauthorizeUser, errors.New("error user unauthorize")))
			return
		}

		c.Next()
	}
}
