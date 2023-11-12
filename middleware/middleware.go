package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if strings.HasPrefix(authHeader, BearerSchema) {
			tokenString := authHeader[len(BearerSchema):]
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if jwt.GetSigningMethod("HS256") != token.Method {
					return nil, jwt.ErrSignatureInvalid
				}
				return []byte("yourSecret"), nil
			})

			if err == nil && token.Valid {
				c.Next()
				return
			}

			c.JSON(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		c.JSON(http.StatusUnauthorized, "unauthorized")
		c.Abort()
	}
}

func GenerateToken(c *gin.Context) {

	claims := jwt.MapClaims{
		"iss": "your-app-name",                       // who creates the token and signs it
		"aud": "your-app-client-id",                  // to whom the token is intended to be sent
		"exp": time.Now().Add(time.Hour * 72).Unix(), // time when the token will expire (72 hours from now)
		"jti": "unique-identifier-for-this-token",
		"iat": time.Now().Unix(), // when the token was issued/created (now)
		"nbf": time.Now().Unix(), // time before which the token is not yet valid (now)
		"sub": "subject",         // the subject/principal is whom the token is about
		"user": map[string]string{ // custom claims
			"id":   "user-id",
			"name": "user-name",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("yourSecret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
