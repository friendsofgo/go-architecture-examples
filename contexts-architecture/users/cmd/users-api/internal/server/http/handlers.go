package http

import (
	"fmt"
	"net/http"

	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/kit/errors"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/users/cmd/users-api/internal/server/http/jwt"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/users/internal/users/creating"
	"github.com/friendsofgo/go-architecture-examples/contexts-architecture/users/internal/users/fetching"

	"github.com/gin-gonic/gin"
)

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
			return
		}

		user, err := fetchService.FetchByID(userID)
		if err != nil {
			if errors.IsNotFound(err) {
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

		user, err := createService.Create(req.Name, req.Email, req.Password)
		if err != nil {
			if errors.IsWrongInput(err) {
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
