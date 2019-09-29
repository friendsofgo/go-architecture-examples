package http

import (
	"net/http"

	"github.com/friendsofgo/go-architecture-examples/hex/internal/creating"
	"github.com/friendsofgo/go-architecture-examples/hex/internal/fetching"
	"github.com/friendsofgo/go-architecture-examples/hex/internal/incrementing"
	"github.com/friendsofgo/go-architecture-examples/hex/internal/server/jwt"

	"github.com/gin-gonic/gin"
)

func MainHandler(
	fetchService fetching.Service,
	createService creating.Service,
	incrementService incrementing.Service,
) (http.Handler, error) {

	r := gin.New()

	// Auth (JWT) handler initialization
	authMiddleware, err := jwt.NewGinMiddleware()
	if err != nil {
		return nil, err
	}

	auth := r.Group("")
	auth.Use(authMiddleware.MiddlewareFunc())

	auth.POST("/counters", createCounterHandlerBuilder(createService))
	auth.GET("/counters/:counterID", getCounterHandlerBuilder(fetchService))
	auth.POST("/counters/increment", incrementCounterHandlerBuilder(fetchService, incrementService))

	return r, nil
}