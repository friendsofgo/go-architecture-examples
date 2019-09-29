package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/friendsofgo/go-architecture-examples/hex/internal/creating"
	"github.com/friendsofgo/go-architecture-examples/hex/internal/errors"
	"github.com/friendsofgo/go-architecture-examples/hex/internal/fetching"
	"github.com/friendsofgo/go-architecture-examples/hex/internal/incrementing"
	"github.com/friendsofgo/go-architecture-examples/hex/internal/server/jwt"
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
			if errors.IsWrongInput(err) {
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
		counter, err := fetchService.Fetch(counterID)
		if err != nil {
			if errors.IsNotFound(err) {
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

		counter, err := fetchService.Fetch(req.ID)
		if err != nil {
			if errors.IsNotFound(err) {
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
			if errors.IsNotFound(err) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}