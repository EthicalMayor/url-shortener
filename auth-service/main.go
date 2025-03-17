package main

import (
	"log"
	"url-shortener/auth-service/database"
	"url-shortener/auth-service/middleware"
	"url-shortener/auth-service/models"
	"url-shortener/auth-service/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		log.Println("Incoming request:", c.Method(), c.Path())
		return c.Next()
	})
	database.ConnectDatabase()

	log.Println("URL Shortener service is running on port 5001") // Log server start

	// Public routes
	app.Post("/register", Register)
	app.Post("/login", Login)

	// Protected route (test middleware)
	app.Get("/protected", middleware.JWTMiddleware(), func(c *fiber.Ctx) error {
		username := c.Locals("username").(string)
		return c.JSON(fiber.Map{"message": "Access granted", "user": username})
	})

	log.Fatal(app.Listen(":5000"))
}

// User registration
func Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	database.DB.Create(&user)
	return c.JSON(fiber.Map{"message": "User registered successfully"})
}

// User login
func Login(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	var user models.User
	database.DB.Where("username = ?", input.Username).First(&user)

	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token, _ := utils.GenerateToken(user.Username)
	return c.JSON(fiber.Map{"token": token})
}
