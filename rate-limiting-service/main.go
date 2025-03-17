package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()
var redisClient *redis.Client

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		log.Println("Incoming request:", c.Method(), c.Path())
		return c.Next()
	})

	log.Println("URL Shortener service is running on port 5001") // Log server start
	// Connect to Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	app.Use(RateLimitMiddleware())

	app.Post("/shorten", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "URL shortened successfully"})
	})

	log.Fatal(app.Listen(":5002"))
}

// Rate Limit Middleware
func RateLimitMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userIP := c.IP()
		key := fmt.Sprintf("rate:%s", userIP)

		// Get the number of requests
		count, _ := redisClient.Get(ctx, key).Int()
		if count >= 5 {
			return c.Status(429).JSON(fiber.Map{"error": "Rate limit exceeded. Try again later."})
		}

		// Increment request count
		redisClient.Incr(ctx, key)
		redisClient.Expire(ctx, key, time.Minute)

		return c.Next()
	}
}
