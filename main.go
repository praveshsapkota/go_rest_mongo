package main

import (
	"os"

	"example/hello/pkg/routes"
	"example/hello/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// app := fiber.New(fiber.Config{
	// 	Prefork:       true,
	// 	CaseSensitive: true,
	// 	StrictRouting: true,
	// 	ServerHeader:  "Fiber",
	// 	AppName:       "test app",
	// })

	// app.Use(cors.New())
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "*",
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// }))

	// api := app.Group("/api") // /api

	// v1 := api.Group("/v1")

	// v1.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("inside")
	// })
	// app.Get("/outside", func(c *fiber.Ctx) error {
	// 	return c.SendString("outside")
	// })
	// app.Get("/in", func(c *fiber.Ctx) error {
	// 	return c.SendString("adsfaf")
	// })
	// println("server connected")

	// log.Fatal(app.Listen(":" + port))

	err := run()
	if err != nil {
		print(err)
		panic(err)
	}

}

func run() error {
	// init env
	err := utils.LoadEnv()
	if err != nil {
		print(err)
		return err
	}

	// init db
	err = utils.InitDB()
	if err != nil {
		return err
	}

	// defer closing db
	defer utils.CloseDB()

	// create app
	app := fiber.New(

		fiber.Config{
			Prefork:       true,
			CaseSensitive: true,
			// StrictRouting: true,
			ServerHeader:  "Fiber",
			AppName:       "test app",
			// JSONEncoder:   json.Marshal,
			// JSONDecoder:   json.Unmarshal,
		})

	// add basic middleware
	app.Use(logger.New())
	app.Use(recover.New())
	// app.Use(cors.New(
	// // 	cors.Config{
	// // 	AllowOrigins: "*",
	// // 	AllowHeaders: "Origin, Content-Type, Accept",	
	// // }
	// ))

	// add routes
	routes.ADDItemGroup(app)

	// start server
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	app.Listen(":"+ port)

	return nil
}
