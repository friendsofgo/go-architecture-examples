package http

import (
	"net/http"

	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/counters/cmd/internal/server/http/jwt"

	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/counters/internal/counters/creating"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/counters/internal/counters/fetching"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/counters/internal/counters/incrementing"

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

	r.Use(gin.Logger(), gin.Recovery())

	auth := r.Group("")
	auth.Use(authMiddleware.MiddlewareFunc())

	auth.POST("/counters", createCounterHandlerBuilder(createService))
	auth.GET("/counters/:counterID", getCounterHandlerBuilder(fetchService))
	auth.POST("/counters/increment", incrementCounterHandlerBuilder(fetchService, incrementService))

	return r, nil
}
