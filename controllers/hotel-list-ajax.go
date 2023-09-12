package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/beego/beego/v2/server/web"
	"travel_guide/models"
)

type HotelControllerAjax struct {
	web.Controller
}

func (c *HotelControllerAjax) Get() {
	// Get form data from the request
	destination := c.GetString("destination")
	checkIn := c.GetString("t-start")
	checkOut := c.GetString("t-end")

	url := "https://booking-com13.p.rapidapi.com/stays/properties/list-v2" +
		"?location=" + destination +
		"&checkin_date=" + checkIn +
		"&checkout_date=" + checkOut +
		"&language_code=en-us&currency_code=USD"

	// Create a channel to receive the result
	resultChan := make(chan interface{}, 1)

	go func() {
		// Create an HTTP request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			resultChan <- map[string]interface{}{"Error": "Error creating request"}
			return
		}

		// Set headers required for the RapidAPI endpoint
		req.Header.Add("X-RapidAPI-Key", "8c45aa97c7msh56eb8adeff6c7fbp15f018jsn948481d80bee")
		req.Header.Add("X-RapidAPI-Host", "booking-com13.p.rapidapi.com")

		// Make the HTTP request
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			resultChan <- map[string]interface{}{"Error": "Error making the request"}
			return
		}
		defer res.Body.Close()

		// Check the HTTP status code
		if res.StatusCode != http.StatusOK {
			resultChan <- map[string]interface{}{"Error": "API request failed with status code " + fmt.Sprint(res.StatusCode)}
			return
		}

		// Read the response body
		body, err := io.ReadAll(res.Body)
		if err != nil {
			resultChan <- map[string]interface{}{"Error": "Error reading the response"}
			return
		}

		// Parse the JSON response into a struct
		var hotelResponse models.HotelInfo
		if err := json.Unmarshal(body, &hotelResponse); err != nil {
			resultChan <- map[string]interface{}{"Error": "Error parsing JSON response"}
			return
		}

		// Send the parsed result through the channel
		resultChan <- hotelResponse
	}()

	// Wait for the result from the goroutine
	result := <-resultChan

	c.Data["json"] = result
	c.ServeJSON()
}
