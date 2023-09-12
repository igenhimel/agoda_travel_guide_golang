package models

type FlightData struct {
	CurrentPage string  `json:"currentPage"`
	   Flight struct {
		 FlightsCount int `json:"filteredFlightsCount"`
	   }
}