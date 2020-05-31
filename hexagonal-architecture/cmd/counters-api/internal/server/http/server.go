package http

import (
	"net/http"

	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/cmd/counters-api/internal/server/jwt"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/creator"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/fetcher"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/incrementer"

	"github.com/gin-gonic/gin"
)

func MainHandler(
	fetchService fetcher.Service,
	createService creator.Service,
	incrementService incrementer.Service,
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
	auth.POST("/counters", createCounterHandlerBuilder(createService))
	auth.GET("/counters/:counterID", getCounterHandlerBuilder(fetchService))
	auth.POST("/counters/increment", incrementCounterHandlerBuilder(fetchService, incrementService))

	return r, nil
}
