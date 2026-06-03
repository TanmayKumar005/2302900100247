package clients

import (
	"encoding/json"
	"net/http"

	"github.com/TanmayKumar005/2302900100247/vehicle_scheduling/internal/models"
)

const BaseURL = "http://4.224.186.213/evaluation-service"

func GetDepots(token string) ([]models.Depot, error) {
	req, _ := http.NewRequest(
		"GET",
		BaseURL+"/depots",
		nil,
	)

	req.Header.Set("Authorization", token)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result struct {
		Depots []models.Depot `json:"depots"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)

	return result.Depots, err
}

func GetVehicles(token string) ([]models.Vehicle, error) {
	req, _ := http.NewRequest(
		"GET",
		BaseURL+"/vehicles",
		nil,
	)
	req.Header.Set("Authorization", token)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result struct {
		Vehicles []models.Vehicle `json:"vehicles"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)

	return result.Vehicles, err
}
