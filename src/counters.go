package main

import (
	"math/rand/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
)

// Define prometheus counters
var (
	TOTAL_RESPONSES = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "hello_world_total",
		Help: "Total number of Hello World responses served",
	})

	TOTAL_REQUESTS = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "hello_world_requests",
		Help: "Total number of Hello World requests received",
	})

	EXCEPTIONS = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "hello_world_exceptions_total",
		Help: "Exceptions serving Hello World.",
	})
)

func helloHandler(c *fiber.Ctx) error {
	// Record the request metric
	TOTAL_REQUESTS.Inc()
	defer TOTAL_RESPONSES.Inc()

	// Gauge of in progress requests
	INPROGRESS.Inc()
	defer INPROGRESS.Dec()

	if rand.Float64() < 0.2 {
		EXCEPTIONS.Inc()
		return fiber.NewError(fiber.StatusInternalServerError, "Random exception occured")
	}

	// Set the last time handler was invoked
	LAST.SetToCurrentTime()

	return c.SendString("Hello, World! This is a Fiber server.")
}
