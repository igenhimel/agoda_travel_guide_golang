package models

type FlightData struct {
	CurrentPage int `json:"currentPage"`
	Flight      struct {
		FlightsCount int `json:"filteredFlightsCount"`
		Data         []struct {
			TravelPrices []struct {
				Price struct {
					Price struct {
						Currency struct {
							Code string `json:"code"`
						} `json:"currency"`
						Value int `json:"value"`
					} `json:"price"`
					Vat struct {
						Value int `json:"value"`
					} `json:"vat"`
				} `json:"price"`
			} `json:"travelerPrices"`
			ShareableUrl string `json:"shareableUrl"`
		} `json:"flights"`
		Routes []struct {
			Destination struct {
				CityCode string `json:"cityCode"`
				CityName string `json:"cityName"`
			} `json:"destination"`
			Origin struct {
				CityCode string `json:"cityCode"`
				CityName string `json:"cityName"`
			} `json:"origin"`
		} `json:"routes"`
	} `json:"data"`
}