package main

import (
	"context"
	"encoding/base64"
	"log"
	"os"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/gofiber/fiber/v2"
)

type HtmlRequest struct {
	HTML     string `json:"html"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
	FullPage bool   `json:"full_page,omitempty"`
	Quality  int    `json:"quality,omitempty"`
}

type HtmlImageResponse struct {
	Success   bool   `json:"success"`
	Message   string `json:"message,omitempty"`
	ImageData string `json:"image_data,omitempty"` // base64 encoded image
}

func main() {
	// Initialize ChromeDP context
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Create ChromeDP context
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Initialize Fiber app
	app := fiber.New()

	// Add middleware for CORS and request logging
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Method() == "OPTIONS" {
			return c.SendStatus(204)
		}

		return c.Next()
	})

	// Add health check endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "healthy",
		})
	})

	// Add HTML to image conversion endpoint
	app.Post("/html-to-image", func(c *fiber.Ctx) error {
		return convertHTMLToImage(c, ctx)
	})

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

func convertHTMLToImage(c *fiber.Ctx, ctx context.Context) error {
	// Parse request body
	var req HtmlRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(HtmlImageResponse{
			Success: false,
			Message: "Invalid request body",
		})
	}

	// Validate HTML content
	if req.HTML == "" {
		return c.Status(fiber.StatusBadRequest).JSON(HtmlImageResponse{
			Success: false,
			Message: "HTML content is required",
		})
	}

	// Set default values
	if req.Width == 0 {
		req.Width = 800
	}
	if req.Height == 0 {
		req.Height = 600
	}
	if req.Quality == 0 {
		req.Quality = 90
	}

	// Create a new context with timeout
	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var buf []byte

	// Run chromedp tasks
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}
			return page.SetDocumentContent(frameTree.Frame.ID, req.HTML).Do(ctx)
		}),
		chromedp.Sleep(1*time.Second), // Wait for rendering
		chromedp.EmulateViewport(int64(req.Width), int64(req.Height)),
		chromedp.CaptureScreenshot(&buf),
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(HtmlImageResponse{
			Success: false,
			Message: "Failed to convert HTML to image: " + err.Error(),
		})
	}

	// Encode image to base64
	base64Image := base64.StdEncoding.EncodeToString(buf)

	return c.JSON(HtmlImageResponse{
		Success:   true,
		ImageData: base64Image,
	})
}
