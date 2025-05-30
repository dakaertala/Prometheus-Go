package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	// Register metrics with prometheus
	prometheus.MustRegister(TOTAL_REQUESTS)
	prometheus.MustRegister(TOTAL_RESPONSES)
	prometheus.MustRegister(EXCEPTIONS)
	prometheus.MustRegister(INPROGRESS)
	prometheus.MustRegister(LAST)
	prometheus.MustRegister(CURRENT_TIME)
}

func main() {
	// Start Prometheus metrics server on port 8000
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println("Metrics server starting on :8000")
		if err := http.ListenAndServe(":8000", nil); err != nil {
			log.Fatal("Metrics server failed:", err)
		}
	}()

	// Create fiber app
	app := fiber.New(fiber.Config{
		ServerHeader: "Fiber",
		AppName:      "My App v1.0.0",
	})

	// Add logger middleware
	app.Use(logger.New())

	// Define routes
	app.Get("/", helloHandler)
	app.Get("/manual", helloHandlerManualTime)

	// You can add more routes here
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	log.Println("Main server starting on localhost:8001")
	log.Println("Metrics available at http://localhost:8000/metrics")

	// Start the Fiber server
	if err := app.Listen(":8001"); err != nil {
		log.Fatal("Main server failed:", err)
	}
}
