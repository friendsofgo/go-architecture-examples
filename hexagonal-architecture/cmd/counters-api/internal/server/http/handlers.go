package http

import (
	"errors"
	"fmt"
	"net/http"

	counters "github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal"
	"github.com/gin-gonic/gin"

	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/cmd/counters-api/internal/server/jwt"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/creating"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/fetching"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/incrementing"
)

func createCounterHandlerBuilder(createService creating.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req CreateCounterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		authorizedUserData, _ := c.Get(jwt.IdentityKey)
		authorizedUser := authorizedUserData.(jwt.User)

		counter, err := createService.CreateCounter(req.Name, authorizedUser.ID)
		if err != nil {
			if errors.Is(err, counters.ErrCreatingUser) {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}

			return
		}

		c.JSON(http.StatusOK, CreateCounterResponse{
			ID:    counter.ID,
			Name:  counter.Name,
			Value: counter.Value,
		})
	}
}

func getCounterHandlerBuilder(
	fetchService fetching.Service,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		counterID := c.Param("counterID")
		counter, err := fetchService.FetchCounterByID(counterID)
		if err != nil {
			if errors.Is(err, counters.ErrCounterNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		authorizedUserData, _ := c.Get(jwt.IdentityKey)
		authorizedUser := authorizedUserData.(jwt.User)
		if authorizedUser.ID != counter.BelongsTo {
			errMsg := fmt.Sprintf("user id %s is not authorized to read the counter %s", authorizedUser.ID, counterID)
			c.JSON(http.StatusForbidden, gin.H{"error": errMsg})
			return
		}

		c.JSON(http.StatusOK, GetCounterResponse{
			ID:    counter.ID,
			Name:  counter.Name,
			Value: counter.Value,
		})
	}
}

func incrementCounterHandlerBuilder(
	fetchService fetching.Service,
	incrementService incrementing.Service,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req IncrementCounterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		counter, err := fetchService.FetchCounterByID(req.ID)
		if err != nil {
			if errors.Is(err, counters.ErrCounterNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		authorizedUserData, _ := c.Get(jwt.IdentityKey)
		authorizedUser := authorizedUserData.(jwt.User)
		if authorizedUser.ID != counter.BelongsTo {
			errMsg := fmt.Sprintf("user id %s is not authorized to increment the counter %s", authorizedUser.ID, req.ID)
			c.JSON(http.StatusForbidden, gin.H{"error": errMsg})
			return
		}

		err = incrementService.Increment(req.ID)
		if err != nil {
			if errors.Is(err, counters.ErrCounterNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}

func getUserHandlerBuilder(
	fetchService fetching.Service,
) func(c *gin.Context) {
	return func(c *gin.Context) {
		authorizedUserData, _ := c.Get(jwt.IdentityKey)
		authorizedUser := authorizedUserData.(jwt.User)

		userID := c.Param("userID")
		if authorizedUser.ID != userID {
			errorMsg := fmt.Sprintf("the user %s cannot read the data of user %s", authorizedUser.ID, userID)
			c.JSON(http.StatusForbidden, gin.H{"error": errorMsg})
		}

		user, err := fetchService.FetchUserByID(userID)
		if err != nil {
			if errors.Is(err, counters.ErrUserNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": err})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			}
			return
		}

		c.JSON(http.StatusOK, GetUserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}
}

func createUserHandlerBuilder(createService creating.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req CreateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := createService.CreateUser(req.Name, req.Email, req.Password)
		if err != nil {
			if  errors.Is(err, counters.ErrCreatingUser) {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}

			return
		}

		c.JSON(http.StatusOK, RegisterUserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}
}
