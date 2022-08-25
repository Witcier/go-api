package middleware

import (
	"strings"

	"witcier/go-api/model/common/response"
	"witcier/go-api/utils"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := StripBearerTokenString(c.Request.Header.Get("Authorization"))
		if token == "" {
			response.Unauthorized(c)
			return
		}

		j := utils.NewJWT()

		claims, err := j.ParseToken(token)
		if err != nil {
			response.Unauthorized(c)
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

func StripBearerTokenString(token string) string {
	// Should be a bearer token
	if len(token) > 6 && strings.ToUpper(token[0:7]) == "BEARER " {
		return token[7:]
	}
	return token
}
