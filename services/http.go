package services

import (
	"aoi/config"
	"aoi/global"
	"aoi/router"
	"encoding/json"
	"log"
	"strings"

	"github.com/antoniodipinto/ikisocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	})

	sites := strings.Join(config.Cors, ", ")
	if sites != "" {
		app.Use(cors.New(cors.Config{
			AllowOrigins: sites,
		}))
	}

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${ip}:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Seoul",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression, // 1
	}))

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

			chat.Clients[id] = kws.UUID
			log.Println("Join root : ", id, "user = ", kws.UUID)
			if _, exists := chat.Rooms[room]; !exists {
				chat.Rooms[room] = make(map[string]string)
			}
			chat.Rooms[room][kws.UUID] = kws.UUID
			log.Println(kws.UUID)

		}))
	}

	app.Static("/webdata", "./webdata")
	app.Static("/", config.DocumentRoot)

	router.SetRouter(app)

	log.Fatal(app.Listen(":" + config.Port))
}
