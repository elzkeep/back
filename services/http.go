package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"path"
	"strings"
	"zkeep/config"
	"zkeep/global"
	"zkeep/router"

	"github.com/antoniodipinto/ikisocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/websocket/v2"
)

func Http() {
	log.Println("HTTP Service Start")

	app := fiber.New(fiber.Config{
		BodyLimit:     500 * 1024 * 1024,
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			// Send custom error page
			err = ctx.Status(code).SendFile(fmt.Sprintf("%v/index.html", config.DocumentRoot))
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			// Return from handler
			return nil
		},
	})

	// 한 요청의 panic으로 서버 전체가 종료되지 않도록 복구
	app.Use(recover.New())

	sites := strings.Join(config.Cors, ", ")
	if sites != "" {
		app.Use(cors.New(cors.Config{
			AllowOrigins: sites,
		}))
	}

	/*
		app.Use(limiter.New(limiter.Config{
			Expiration: 1 * time.Second,
			Max:        2,
		}))
	*/

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${ip}:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Seoul",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression, // 1
	}))

	/*
		app.Get("/*", func(ctx *fiber.Ctx) error {
			return ctx.SendFile(fmt.Sprintf("%v/index.html", config.DocumentRoot), true)
		})
	*/

	if chat.Use == true {
		app.Use("/ws", func(c *fiber.Ctx) error {
			if websocket.IsWebSocketUpgrade(c) {
				log.Println("upgrade")
				return c.Next()
			}

			return fiber.ErrUpgradeRequired
		})

		app.Get("/ws/:id", ikisocket.New(func(kws *ikisocket.Websocket) {
			id := kws.Params("id")
			kws.SetAttribute("id", id)

			room := global.Atol(id)

			chat.Join(id, room, kws.UUID)
			log.Println("Join root : ", id, "user = ", kws.UUID)

		}))
	}

	app.Get("/.well-known/pki-validation/0AFE49F541E6BB3A544AE12E290DD8DC.txt", func(c *fiber.Ctx) error {
		return c.SendString("EF1A2C0658FF86E54CF0E17A385ADADF29B1E9FFDEC29274984883A791A2BB23\nsectigo.com\ndcv20231207d7519")
	})

	app.Static("/webdata", "./webdata")
	app.Static("/", config.DocumentRoot)
	//app.Static("/assets", fmt.Sprintf("%v/assets", config.DocumentRoot))

	/*
		app.Get("/*", func(ctx *fiber.Ctx) error {
			return ctx.SendFile(fmt.Sprintf("%v/index.html", config.DocumentRoot))
		})
	*/

	// 응답 비밀번호 제거·수정 시 보존 → 회원가입(공개) → 인증·업체 분리가 걸린 API 순
	app.Use(router.SanitizeResponse)
	app.Use(router.KeepSecretsOnUpdate)
	router.SetSignupRouter(app)
	router.SetRouter(app)

	// SPA 화면 경로 새로고침 시 index.html 반환
	app.Get("/*", func(c *fiber.Ctx) error {
		if strings.HasPrefix(c.Path(), "/api/") {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.SendFile(path.Join(config.DocumentRoot, "index.html"))
	})

	log.Fatal(app.Listen(":" + config.Port))
}
