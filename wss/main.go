package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		// prefork: true,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/user/login", websocket.New(UserLogin))

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		log.Println(c.Locals("allowed"))  // true
		log.Println(c.Params("id"))       // 123
		log.Println(c.Query("v"))         // 1.0
		log.Println(c.Cookies("session")) // ""

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			mt  int
			msg []byte
			err error
		)

		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))

	log.Fatal(app.Listen(":3000"))
	// Access the websocket server: ws://localhost:3000/ws/123?v=1.0
	// https://www.websocket.org/echo.html
}

func UserLogin(c *websocket.Conn) {

	var Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	for {
		if err := c.ReadJSON(&Request); err != nil {
			log.Println("client connection closed")
			break
		}

		fmt.Println(Request)

		res := fiber.Map{"message": "", "user_info_id": 0}
		email := strings.ToLower(strings.TrimSpace(Request.Email))

		if email == "test@naver.com" {
			res = fiber.Map{
				"message":      "hello " + email,
				"user_info_id": 1234,
			}
		}

		if err := c.WriteJSON(&res); err != nil {
			log.Println("client connection closed")
			break
		}

	}
}
