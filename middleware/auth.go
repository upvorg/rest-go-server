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

type contextKey string

const (
	ctxAuthKey     contextKey = "ctxAuthKey"
	JWT_TOKEN_SIGN string     = "JWT_TOKEN_SIGN"
	GinCtxAuthKey  string     = "ginCtxAuthKey"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if !strings.Contains(token, "Bearer") {
			mapClaims := jwt.MapClaims{"user_name": "", "user_id": 0, "authorized": false}
			c.Set(GinCtxAuthKey, mapClaims)
			// ctx := context.WithValue(c.Request.Context(), ctxAuthKey, mapClaims)
			// next.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
			return
		}

		authHeader := strings.Split(token, "Bearer ")
		fmt.Printf("authheader -> %s and len -> %d\n", authHeader, len(authHeader))
		if len(authHeader) != 2 || authHeader[0] == "null" {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Writer.Write([]byte("Malformed Token"))
			log.Fatal("Malformed token")
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(JWT_TOKEN_SIGN), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				// ctx := context.WithValue(c.Request.Context(), ctxAuthKey, claims)
				// Access context values in handlers like this
				//props, _ := c.Request.Context().Value("props").(jwt.MapClaims)

				c.Set(GinCtxAuthKey, claims)
				// next.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
			} else {
				fmt.Println("token err -> ", err)
				c.Writer.WriteHeader(http.StatusUnauthorized)
				c.Writer.Write([]byte("you are Unauthorized or your token is expired"))
			}
		}

	}
}

func CreateToken(userId uint64, name string) (string, error) {
	atClaims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userId,
		"user_name":  name,
		"exp":        time.Now().Add(time.Minute * 15).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(config.JwtSlat))
	if err != nil {
		return "", errors.New("an error occured during the create token")
	}
	fmt.Println("jwt map --> ", atClaims)
	return token, nil
}

func ParseMapClaims(myMap jwt.MapClaims, tokenStr string) jwt.Claims {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_TOKEN_SIGN), nil
	})
	if err != nil {
		log.Fatal("an error occured during the parse jwt ,,,")
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims
}
