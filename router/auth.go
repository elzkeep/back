package router

import (
	"errors"
	"log"
	"net/url"
	"time"
	"zkeep/models"
	"zkeep/models/user"

	"github.com/gofiber/fiber/v2"

	"github.com/dgrijalva/jwt-go/v4"
)

type AuthTokenClaims struct {
	User               models.User `json:"user"`
	jwt.StandardClaims             // 표준 토큰 Claims
}

var _secretCode string = "WkaQHd100%"

var JwtAuthRequired = func(c *fiber.Ctx) error {
	var token string

	if values := c.Get("Authorization"); len(values) > 0 {
		str := values

		if len(str) > 7 && str[:7] == "Bearer " {
			token = str[7:]

			claims := AuthTokenClaims{}
			key := func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("Unexpected Signing Method")
				}
				return []byte(_secretCode), nil
			}

			_, err := jwt.ParseWithClaims(token, &claims, key)
			if err == nil {
				c.Locals("user", &(claims.User))
				return c.Next()
			}
		}
	}

	path := c.Path()
	u, _ := url.Parse(path)

	if c.Method() == "GET" && u.Path == "/api/user" {
		return c.Next()
	}

	if c.Method() == "GET" && u.Path == "/api/webnotice" {
		return c.Next()
	}

	if u.Path == "/api/webquestion" {
		return c.Next()
	}

	if c.Method() == "GET" && u.Path == "/api/web" {
		return c.Next()
	}

	if u.Path == "/api/webjoin" {
		return c.Next()
	}

	if c.Method() == "POST" && u.Path == "/api/user/findid" {
		return c.Next()
	}

	if c.Method() == "POST" && u.Path == "/api/user/password" {
		return c.Next()
	}

	if c.Method() == "POST" && u.Path == "/api/user" {
		return c.Next()
	}

	if c.Method() == "POST" && u.Path == "/api/company" {
		return c.Next()
	}

	if c.Method() == "GET" && u.Path == "/api/company" {
		return c.Next()
	}

	if u.Path == "/api/jwt" {
		return c.Next()
	}

	log.Println("Jwt header is broken")

	return nil
}

func JwtAuth(c *fiber.Ctx, loginid string, passwd string) map[string]interface{} {
	conn := models.NewConnection()

	manager := models.NewUserManager(conn)
	item := manager.GetByLoginid(loginid)

	if item == nil {
		return map[string]interface{}{
			"code":    "error",
			"message": "user not found",
		}
	}

	if item.Passwd != passwd {
		return map[string]interface{}{
			"code":    "error",
			"message": "wrong password",
		}
	}

	if item.Status != user.StatusUse {
		return map[string]interface{}{
			"code":    "error",
			"message": "status error",
		}
	}

	at := AuthTokenClaims{
		User: *item,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * 24 * 365 * 10)),
		},
	}

	atoken := jwt.NewWithClaims(jwt.SigningMethodHS256, &at)
	signedAuthToken, _ := atoken.SignedString([]byte(_secretCode))

	item.Passwd = ""
	return map[string]interface{}{
		"code":  "ok",
		"token": signedAuthToken,
		"user":  item,
	}
}
