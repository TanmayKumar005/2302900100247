package models

type Depot struct {
	ID            int `json:"ID"`
	MechanicHours int `json:"MechanicHoours"`
}

type Vehicle struct {
	TaskID   string `json:"TackID"`
	Duration int    `json:"Duration"`
	Impact   int    `json:"Impact"`
}

type ScheduleResponse struct {
	DepotID       int      `json:"DepotID"`
	SelectedTasks []string `json:"SelectedTasks"`
	TotalDuration int      `json:"TotalDuration"`
	TotalImpact   int      `json:"TotalImpact"`
}
