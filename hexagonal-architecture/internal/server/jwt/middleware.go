package jwt

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/friendsofgo/go-architecture-examples/hexagonal-architecture/internal/fetching"
)

const (
	realmName   = "REALM NAME IS PENDING TBD"
	IdentityKey = "auth"
	authKey     = "friendsofgo"
)

type User struct {
	ID    string
	Name  string
	Email string
}


// NewGinMiddleware returns a JWT middleware for Gin
func NewGinMiddleware(fetchService fetching.Service) (*jwt.GinJWTMiddleware, error) {
	key := []byte(authKey)

	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           realmName,
		Key:             key,
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     IdentityKey,
		PayloadFunc:     payloadHandler,
		IdentityHandler: identityHandler,
		Authenticator:   authHandlerBuilder(fetchService),
		Unauthorized:    unauthorizedHandler,
	})
}

func payloadHandler(data interface{}) jwt.MapClaims {
	if user, ok := data.(*User); ok {
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
	return User{
		ID:    claims["id"].(string),
		Name:  claims["name"].(string),
		Email: claims["email"].(string),
	}
}

type LoginRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func authHandlerBuilder(fetchService fetching.Service) func(c *gin.Context) (interface{}, error) {
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

		return &User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}, nil
	}
}

func unauthorizedHandler(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
