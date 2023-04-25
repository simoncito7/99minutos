package authentication

import (
	"errors"
	"net/http"
	"strings"

	"github.com/99minutos/token"
	"github.com/gin-gonic/gin"
)

const (
	authHeaderKey  = "authorization"
	AuthPayloadKey = "authorization_payload"
)

var (
	_errAuthHeader                   = errors.New("need authorization header to be authenticated")
	_errInvalidAuthHeaderFormat      = errors.New("invalid authorization format")
	_errInvalidAuthHeaderUnsupported = errors.New("authorization type not supported")
)

func AuthMiddelware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(authHeaderKey)
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, _errAuthHeader)
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, _errInvalidAuthHeaderFormat)
			return
		}

		if strings.ToLower(fields[0]) != "bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, _errInvalidAuthHeaderUnsupported)
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		ctx.Set(AuthPayloadKey, payload)
		ctx.Next()
	}
}
