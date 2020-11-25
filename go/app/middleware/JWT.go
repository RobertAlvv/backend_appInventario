package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func JwtMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}

func SignToken(tokenKey, id string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["sub"] = id

	t, err := token.SignedString([]byte(tokenKey))

	if err != nil {
		return ""
	}

	return t
}

func ExtractUserIDFromJWT(bearer, tokenKey string) string {
	//Bearer ey.....
	token := bearer[7:]
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})
	if err != nil {
		return ""
	}

	if t.Valid {
		claims := t.Claims.(jwt.MapClaims)
		return claims["sub"].(string)
	}
	return ""
}
