package servers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AnalyticsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tknStr, err := c.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				c.Abort()
				c.Writer.WriteHeader(http.StatusUnauthorized)
				return
			}

			c.Abort()
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}

		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.Abort()
				c.Writer.WriteHeader(http.StatusUnauthorized)
				return
			}
			c.Abort()
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			c.Abort()
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
