package middleware

import (
	"net/http"
	"witcier/go-api/config"
	"witcier/go-api/global"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token, X-Token, X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT, PATCH")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func CorsByRules() gin.HandlerFunc {
	if global.Config.Cors.Mode == "allow-all" {
		return Cors()
	}

	return func(c *gin.Context) {
		whitelist := checkCors(c.GetHeader("Origin"))

		if whitelist != nil {
			c.Header("Access-Control-Allow-Origin", whitelist.AllowOrigin)
			c.Header("Access-Control-Allow-Headers", whitelist.AllowHeaders)
			c.Header("Access-Control-Allow-Methods", whitelist.AllowMethods)
			c.Header("Access-Control-Expose-Headers", whitelist.ExposeHeaders)
			if whitelist.AllowCredentials {
				c.Header("Access-Control-Allow-Credentials", "true")
			}
		}

		if whitelist == nil && global.Config.Cors.Mode == "strict-whitelist" && !(c.Request.Method == "GET" && c.Request.URL.Path == "/health") {
			c.AbortWithStatus(http.StatusForbidden)
		} else {
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(http.StatusNoContent)
			}
		}

		c.Next()
	}
}

func checkCors(currentOrigin string) *config.CORSWhitelist {
	for _, whitelist := range global.Config.Cors.Whitelist {
		if currentOrigin == whitelist.AllowOrigin {
			return &whitelist
		}
	}

	return nil
}
