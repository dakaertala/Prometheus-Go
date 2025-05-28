package main

import (
	"math/rand/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
)

// Define prometheus counters
var (
	helloWorldsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "hello_world_total",
		Help: "Total number of Hello World responses served",
	})

	helloWorldsRequested = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "hello_world_requests",
		Help: "Total number of Hello World requests received",
	})

	exceptions = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "hello_world_exceptions_total",
		Help: "Exceptions serving Hello World.",
	})
)

func init() {
	// Register metrics with prometheus
	prometheus.MustRegister(helloWorldsRequested)
	prometheus.MustRegister(helloWorldsTotal)
	prometheus.MustRegister(exceptions)
}

func helloHandler(c *fiber.Ctx) error {
	helloWorldsRequested.Inc()

	if rand.Float64() < 0.2 {
		exceptions.Inc()
		return fiber.NewError(fiber.StatusInternalServerError, "Random exception occured")
	}

	defer helloWorldsTotal.Inc()

	return c.SendString("Hello, World! This is a Fiber server.")
}
