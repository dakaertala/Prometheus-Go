package main

import (
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
)

func init() {
	// Register metrics with prometheus
	prometheus.MustRegister(helloWorldsRequested)
	prometheus.MustRegister(helloWorldsTotal)
}

func helloHandler(c *fiber.Ctx) error {
	helloWorldsRequested.Inc()

	defer helloWorldsTotal.Inc()

	return c.SendString("Hello, World! This is a Fiber server.")
}
