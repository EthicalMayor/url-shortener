package main

import (
	"log"
	"math/rand"
	"time"

	"url-shortener/url-service/database"
	"url-shortener/url-service/models"

	"github.com/gofiber/fiber/v2"
)

// Global random generator
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	app := fiber.New() // Initialize Fiber app

	app.Use(func(c *fiber.Ctx) error {
		log.Println("Incoming request:", c.Method(), c.Path())
		return c.Next()
	}) //	Logging every request

	database.ConnectDatabase() // Connect to the database

	log.Println("URL Shortener service is running on port 5001") // Log server start

	app.Post("/shorten", ShortenURL)   // Route to shorten a URL
	app.Get("/:shortCode", ResolveURL) // Route to resolve and redirect

	log.Fatal(app.Listen(":5001")) // Start the server on port 5001
}

// Generate a unique short code
func generateShortCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortCode := make([]byte, length)

	for i := range shortCode {
		shortCode[i] = charset[rng.Intn(len(charset))]
	}

	return string(shortCode)
}

// Shorten a URL
func ShortenURL(c *fiber.Ctx) error {
	type Request struct {
		OriginalURL string `json:"original_url"`
	}

	var request Request
	if err := c.BodyParser(&request); err != nil {
		log.Println("Error parsing request:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	log.Println("🔹 Received request to shorten URL:", request.OriginalURL)

	// Ensure unique short code
	var shortCode string
	for {
		shortCode = generateShortCode(6)
		var existing models.URL
		database.DB.Where("short_code = ?", shortCode).First(&existing)
		if existing.ID == 0 { // Unique code found
			break
		}
	}

	log.Println("Generated unique short code:", shortCode)

	url := models.URL{ShortCode: shortCode, Original: request.OriginalURL}
	if err := database.DB.Create(&url).Error; err != nil {
		log.Println("Error saving URL to database:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save URL"})
	}

	shortenedURL := "http://localhost:5001/" + shortCode
	log.Println("Successfully shortened:", request.OriginalURL, "➡", shortenedURL)

	return c.JSON(fiber.Map{"short_url": shortenedURL})
}

// Retrieve the original URL
func ResolveURL(c *fiber.Ctx) error {
	shortCode := c.Params("shortCode")
	var url models.URL

	if err := database.DB.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "URL not found"})
	}

	return c.Redirect(url.Original, 301) // 301 for permanent redirect
}
