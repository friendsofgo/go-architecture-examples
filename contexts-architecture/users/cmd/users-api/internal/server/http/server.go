package http

import (
	"net/http"

	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/users/cmd/users-api/internal/server/http/jwt"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/users/internal/users/creating"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/users/internal/users/fetching"

	"github.com/gin-gonic/gin"
)

func MainHandler(
	fetchService fetching.Service,
	createService creating.Service,
) (http.Handler, error) {

	r := gin.New()

	// Auth (JWT) handler initialization
	authMiddleware, err := jwt.NewGinMiddleware(fetchService)
	if err != nil {
		return nil, err
	}

	r.Use(gin.Logger(), gin.Recovery())

	r.POST("/users/register", createUserHandlerBuilder(createService))
	r.POST("/users/login", authMiddleware.LoginHandler)

	auth := r.Group("")
	auth.Use(authMiddleware.MiddlewareFunc())

	auth.GET("/users/:userID", getUserHandlerBuilder(fetchService))

	return r, nil
}
