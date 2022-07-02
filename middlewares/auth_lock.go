package middlewares

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	api "github.com/nsocrates/go-blog/api/common"
	"github.com/nsocrates/go-blog/common"
)

func AuthLock(isLocked bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		api.UpdateContextUserModel(c, 0)
		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(common.SECRET))
			return b, nil
		})

		if err != nil {
			if isLocked {
				c.AbortWithError(http.StatusUnauthorized, err)
			}

			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			myUserId := uint(claims["id"].(float64))
			api.UpdateContextUserModel(c, myUserId)
		}
	}
}
