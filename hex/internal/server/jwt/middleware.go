package jwt

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
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
func NewGinMiddleware() (*jwt.GinJWTMiddleware, error) {
	key := []byte(authKey)

	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           realmName,
		Key:             key,
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     IdentityKey,
		IdentityHandler: identityHandler,
		Unauthorized:    unauthorizedHandler,
	})
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return User{
		ID:    claims["id"].(string),
		Name:  claims["name"].(string),
		Email: claims["email"].(string),
	}
}

func unauthorizedHandler(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
