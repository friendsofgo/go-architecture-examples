package jwt

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	counters "github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal"
	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/platform/server/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/fetching"
)

// NewMiddleware returns a JWT middleware for Gin
func NewMiddleware(fetchService fetching.Service) (*jwt.GinJWTMiddleware, error) {
	key := []byte(http.AuthKey)

	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           http.RealmName,
		Key:             key,
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     http.IdentityKey,
		PayloadFunc:     payloadHandler,
		IdentityHandler: identityHandler,
		Authenticator:   auth(fetchService),
		Unauthorized:    unauthorized,
	})
}

func payloadHandler(data interface{}) jwt.MapClaims {
	if user, ok := data.(*counters.User); ok {
		return jwt.MapClaims{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		}
	}
	return jwt.MapClaims{}
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return counters.User{
		ID:    claims["id"].(string),
		Name:  claims["name"].(string),
		Email: claims["email"].(string),
	}
}

type LoginRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func auth(fetchService fetching.Service) func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (i interface{}, e error) {
		var req LoginRequest
		if err := c.ShouldBind(&req); err != nil {
			return "", jwt.ErrMissingLoginValues
		}

		user, err := fetchService.FetchUserByEmail(req.Email)
		if err != nil {
			return nil, jwt.ErrFailedAuthentication
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password))
		if err != nil {
			return nil, jwt.ErrFailedAuthentication
		}

		return &counters.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}, nil
	}
}

func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
