package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/TanmayKumar005/2302900100247/vehicle_scheduling/internal/clients"
	"github.com/TanmayKumar005/2302900100247/vehicle_scheduling/internal/services"
)

func ScheduleVehicles(c *gin.Context) {
	token := c.GetHeader("Authorization")

	depots, err := clients.GetDepots(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	vehicles, err := clients.GetVehicles(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	depot := depots[0]

	selected, TotalImpact := services.OptimizeTasks(vehicles, depot.MechanicHours)

	totalDuration := 0
	var taskIDs []string

	for _, v := range selected {
		totalDuration += v.Duration
		taskIDs = append(taskIDs, v.TaskID)
	}

	c.JSON(http.StatusOK, gin.H{
		"DepotID":       depot.ID,
		"SelectedTasks": taskIDs,
		"TotalDuration": totalDuration,
		"TotalImpact":   TotalImpact,
	})
}
