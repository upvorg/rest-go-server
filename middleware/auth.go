package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"upv.life/server/config"
)

const (
	CTX_AUTH_KEY string = "ginCtxAuthKey"
)

var JWT_TOKEN_SIGN string = config.JwtSalt

type AuthClaims struct {
	jwt.StandardClaims
	Name   string `json:"name,omitempty"`
	UserId uint64 `json:"uid,omitempty"`
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if !strings.Contains(token, "Bearer") {
			return
		}

		authHeader := strings.Split(token, "Bearer ")
		if len(authHeader) != 2 || authHeader[0] == "null" {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Malformed token",
			})
			log.Fatal("Malformed token")
		} else {
			if claims, err := ParseJwtToken(authHeader[1]); claims != nil {
				// http router 写法
				// ctx := context.WithValue(c.Request.Context(), ctxAuthKey, claims)
				// next.ServeHTTP(c.Writer, c.Request.WithContext(ctx))

				// Access context values in handlers like this
				// props, _ := c.Request.Context().Value("props").(jwt.MapClaims)

				c.Set(CTX_AUTH_KEY, claims)
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
				c.JSON(http.StatusUnauthorized, gin.H{
					"msg": err.Error(),
				})
			}

		}

	}
}

func GenerateJwtToken(userId uint64, name string) (string, error) {
	authClaims := AuthClaims{
		Name:   name,
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			Id:        fmt.Sprintf("%d", userId),
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims).SignedString([]byte(JWT_TOKEN_SIGN))
	return token, err
}

func ParseJwtToken(tokenStr string) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWT_TOKEN_SIGN), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				log.Fatal(errors.New("that's not even a token"))
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				log.Fatal(errors.New("token is expired"))
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				log.Fatal(errors.New("token not active yet"))
				return nil, errors.New("token not active yet")
			} else {
				log.Fatal(errors.New("couldn't handle this token"))
				return nil, errors.New("couldn't handle this token")
			}
		}
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("unknow token")
}
