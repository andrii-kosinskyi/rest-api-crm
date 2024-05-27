package model

type Client struct {
	ID                int64  `json:"id"`
	Name              string `json:"name"`
	WorkingHoursStart string `json:"working_hours_start"`
	WorkingHoursEnd   string `json:"working_hours_end"`
	Priority          int64  `json:"priority"`
	LeadCapacity      int64  `json:"lead_capacity"`
}
