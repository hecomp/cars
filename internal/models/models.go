package models

// Car represents a Car part of a Car Request
// swagger:model
type Car struct {
	Id       string `json:"id"`
	Make     string `json:"make"`
	Model    string `json:"model"`
	Package  string `json:"package"`
	Color    string `json:"color"`
	Year     int    `json:"year"`
	Category string `json:"Category"`
	Mileage  int    `json:"mileage"`
	Price    int    `json:"price"`
}

// HealthResponse contains the current status of the application instance.
type HealthResponse struct {
	Status string `json:"status"`
}
