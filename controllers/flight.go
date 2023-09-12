package controllers

import (
	"encoding/json"
	"net/http"
    "travel_guide/models"
	"github.com/beego/beego/v2/server/web"
)

type FlightController struct {
	web.Controller
}

func fetchFlightData(url string, ch chan<- models.FlightData) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ch <- models.FlightData{} // Send an empty FlightData on error
		return
	}

	req.Header.Add("X-RapidAPI-Key", "8c45aa97c7msh56eb8adeff6c7fbp15f018jsn948481d80bee")
	req.Header.Add("X-RapidAPI-Host", "booking-com13.p.rapidapi.com")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		ch <- models.FlightData{} // Send an empty FlightData on error
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		ch <- models.FlightData{} // Send an empty FlightData on non-OK status code
		return
	}

	var responseData models.FlightData

	if err := json.NewDecoder(res.Body).Decode(&responseData); err != nil {
		ch <- models.FlightData{} // Send an empty FlightData on JSON parsing error
		return
	}

	ch <- responseData
}

func (c *FlightController) Get() {
   
	source := c.GetString("location_from");
	destination:= c.GetString("location_to");
	departure_date := c.GetString("departure_date");
	class := c.GetString("class")
	return_date := c.GetString("return_date");

	if source == "" || destination == "" || departure_date == "" || class == "" {
		GlobalError = "Please Fill the all Required Field"
        c.Redirect("/", 302)
        return
	}

   
    var url string;

   if return_date == "" {
    url = "https://booking-com13.p.rapidapi.com/flights/one-way?" +
        "location_from=" + source +
        "&location_to=" + destination +
        "&departure_date=" + departure_date +
        "&page=1&country_flag=us" +
        "&class=" + class
   }else {
    url = "https://booking-com13.p.rapidapi.com/flights/return?" +
        "location_from=" + source +
        "&location_to=" + destination +
        "&departure_date=" + departure_date +
        "&return_date=" + return_date +
        "&page=1&country_flag=us" +
        "&class=" + class
}


	// Create a channel to receive flight data
	ch := make(chan models.FlightData)

	// Start a goroutine to fetch flight data
	go fetchFlightData(url, ch)

	// Process the received flight data
	responseData := <-ch

	if responseData.CurrentPage > 0 {
		c.Data["CurrentPage"] = responseData.CurrentPage
		c.Data["Flight"] = responseData.Flight
		c.Data["travel_date"]=departure_date
		if return_date!=""{
			c.Data["return_date"]=return_date
	} else {
		// Handle the case where the request failed
		// You can set appropriate error data or redirect to an error page
	}

	// Render the "flight.tpl" template
	c.TplName = "flight/flight.tpl"
	c.Render()
}
}
