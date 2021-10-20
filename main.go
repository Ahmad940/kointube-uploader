package main

import (
	"context"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	var cld, err = cloudinary.New()
	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
	}

	ctx := context.Background()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	// Routes
	app.Post("/upload/video", func(c *fiber.Ctx) error {
		// Get first file from form field "document":
		file, err := c.FormFile("video")
		if err != nil {
			return err
		}

		tempFile, err := file.Open()
		if err != nil {
			return err
		}

		id, err := gonanoid.New()
		if err != nil {
			return err
		}

		uploadResult, err := cld.Upload.Upload(
			ctx,
			tempFile,
			uploader.UploadParams{PublicID: id, Folder: "kointube/videos"})
		if err != nil {
			log.Fatalf("Failed to upload file, %v\n", err)
		}

		return c.JSON(map[string]string{
			"SecureURL": uploadResult.SecureURL,
		})
	})

	// Routes
	app.Post("/upload/image", func(c *fiber.Ctx) error {
		// Get first file from form field "document":
		file, err := c.FormFile("image")
		if err != nil {
			return err
		}

		tempFile, err := file.Open()
		if err != nil {
			return err
		}

		id, err := gonanoid.New()
		if err != nil {
			return err
		}

		uploadResult, err := cld.Upload.Upload(
			ctx,
			tempFile,
			uploader.UploadParams{PublicID: id, Folder: "kointube/images"})
		if err != nil {
			log.Fatalf("Failed to upload file, %v\n", err)
		}

		return c.JSON(map[string]string{
			"SecureURL": uploadResult.SecureURL,
		})
	})

	log.Fatal(app.Listen(os.Getenv("PORT")))
}
