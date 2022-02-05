package middleware

import (
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
	GinCtxAuthKey string = "ginCtxAuthKey"
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
			// mapClaims := AuthClaims{}
			// c.Set(GinCtxAuthKey, mapClaims)
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
			if claims := ParseJwtToken(authHeader[1]); claims != nil {
				// http router 写法
				// ctx := context.WithValue(c.Request.Context(), ctxAuthKey, claims)
				// next.ServeHTTP(c.Writer, c.Request.WithContext(ctx))

				// Access context values in handlers like this
				// props, _ := c.Request.Context().Value("props").(jwt.MapClaims)

				c.Set(GinCtxAuthKey, claims)
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
				c.JSON(http.StatusUnauthorized, gin.H{
					"msg": "You are Unauthorized or your token is expired",
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

func ParseJwtToken(tokenStr string) jwt.Claims {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWT_TOKEN_SIGN), nil
	})
	if err != nil {
		log.Fatal("error while parsing the token" + err.Error())
	}
	if claims, ok := token.Claims.(AuthClaims); ok && token.Valid {
		return claims
	}
	return nil
}
