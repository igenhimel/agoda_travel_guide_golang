package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type FlightController struct {
	web.Controller
}

type FlightData struct {
	CurrentPage int `json:"currentPage"`
	Flight      struct {
		FlightsCount int `json:"filteredFlightsCount"`
		Data         []struct {
			TravelPrices []struct {
				Price struct {
					Price struct {
						Currency struct {
							Code  string `json:"code"`
						} `json:"currency"`
						Value string `json:"value"`
						Vat struct {
							Value string `json:"value"`
						} `json:"vat"`
					} `json:"price"`
				} `json:"price"`
			} `json:"travelerPrices"`
		} `json:"flights"`
	} `json:"data"`
}

func (c *FlightController) Get() {
	url := "https://booking-com13.p.rapidapi.com/flights/one-way?location_from=Dhaka%2C%20Bangladesh&location_to=New%20Delhi%2C%20India&departure_date=2023-09-15&page=1&country_flag=us&class=Business"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Add("X-RapidAPI-Key", "3ad041d245msh9e23b1e8792896cp1fbc15jsn5e20851b7ca7")
	req.Header.Add("X-RapidAPI-Host", "booking-com13.p.rapidapi.com")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("API request failed with status code: %d\n", res.StatusCode)
		return
	}

	var responseData FlightData

	// Parse the JSON response into the responseData struct
	if err := json.NewDecoder(res.Body).Decode(&responseData); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Access and print the structure data
	fmt.Printf("Current Page: %d\n", responseData.CurrentPage)
	fmt.Printf("Filtered Flights Count: %d\n", responseData.Flight.FlightsCount)

	// Print data for each flight
	for i, flight := range responseData.Flight.Data {
		fmt.Printf("Flight %d:\n", i+1)
		// Access and print data for traveler prices in each flight
		for j, travelerPrice := range flight.TravelPrices {
			fmt.Printf("Traveler Price %d:\n", j+1)
			fmt.Printf("Currency Code: %s\n", travelerPrice.Price.Price.Currency.Code)
			fmt.Printf("Price Value: %s\n", travelerPrice.Price.Price.Value)
			fmt.Printf("VAT Value: %s\n", travelerPrice.Price.Price.Vat.Value)
		}
	}
}
