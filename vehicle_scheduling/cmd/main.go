package main

import (
	"github.com/TanmayKumar005/2302900100247/vehicle_scheduling/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET(
		"/api/v1/schedule",
		handlers.ScheduleVehicles,
	)

	r.Run(":2020")
}
