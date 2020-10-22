package users

import (
	"fmt"
	"net/http"

	"github.com/friendsofgo/errors"
	counters "github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/creating"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/fetching"
	"github.com/gin-gonic/gin"
)

func GetUserHandlerBuilder(
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

func CreateUserHandlerBuilder(createService creating.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		var req CreateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := createService.CreateUser(req.Name, req.Email, req.Password)
		if err != nil {
			if errors.Is(err, counters.ErrCreatingUser) {
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
