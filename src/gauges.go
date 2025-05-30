package main

import (
	"math/rand/v2"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	INPROGRESS = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "hello_worlds_inprogress",
		Help: "Number of Hello Worlds in progress.",
	})

	LAST = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "hello_world_last_time_seconds",
		Help: "The last time a Hello World was served.",
	})

	CURRENT_TIME = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "current_time_seconds",
		Help: "The current time",
	}, func() float64 {
		return float64(time.Now().Unix())
	})
)

func helloHandlerManualTime(c *fiber.Ctx) error {
	TOTAL_REQUESTS.Inc()

	INPROGRESS.Inc()
	defer INPROGRESS.Dec()

	if rand.Float64() < 0.2 {
		EXCEPTIONS.Inc()
		return fiber.NewError(fiber.StatusInternalServerError, "Random exception occurred")
	}

	// Manual time setting
	LAST.Set(float64(time.Now().Unix()))

	return c.SendString("Hello world (manual time)")
}
