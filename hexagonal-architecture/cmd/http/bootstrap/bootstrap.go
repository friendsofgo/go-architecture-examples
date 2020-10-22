package bootstrap

import (
	"fmt"
	"log"

	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/creating"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/fetching"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/incrementing"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/platform/server/http"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/platform/server/http/handler/counters"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/platform/server/http/handler/users"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/platform/server/http/middleware/jwt"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/platform/storage/inmemory"
	"github.com/gin-gonic/gin"
)

const (
	ApiHostDefault = "localhost"
	ApiPortDefault = 3000
)

func Run() error {
	var (
		countersRepository = inmemory.NewCountersRepository()
		usersRepository    = inmemory.NewUsersRepository()

		fetchingService     = fetching.NewService(countersRepository, usersRepository)
		creatingService     = creating.NewService(countersRepository, usersRepository)
		incrementingService = incrementing.NewService(countersRepository)
	)

	return runServer(
		fetchingService,
		creatingService,
		incrementingService,
	)
}

func runServer(
	fetchingService fetching.Service,
	creatingService creating.Service,
	incrementingService incrementing.Service,
) error {
	httpAddr := fmt.Sprintf("%s:%d", ApiHostDefault, ApiPortDefault)
	srv := http.NewServer(httpAddr)

	// Auth (JWT) handler initialization
	authMiddleware, err := jwt.NewMiddleware(fetchingService)
	if err != nil {
		return err
	}

	srv.Use(gin.Logger(), gin.Recovery())

	srv.POST("/users/register", users.CreateUser(creatingService))
	srv.POST("/users/login", authMiddleware.LoginHandler)

	auth := srv.Group("")
	auth.Use(authMiddleware.MiddlewareFunc())

	auth.GET("/users/:userID", users.GetUser(fetchingService))
	auth.POST("/counters", counters.CreateCounter(creatingService))
	auth.GET("/counters/:counterID", counters.GetCounter(fetchingService))
	auth.POST("/counters/increment", counters.IncrementCounter(fetchingService, incrementingService))

	log.Println("Server on tap:", httpAddr)
	return srv.Run()
}
