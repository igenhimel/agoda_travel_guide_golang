package controllers

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    beego "github.com/beego/beego/v2/server/web"
	"travel_guide/models"
)

type HotelController struct {
    beego.Controller
}

type HotelData struct {
    ImgURLs []string
    Titles  []string
    Prices  []string
}

func (c *HotelController) Get() {
    // Get form data from the request
    destination := c.GetString("destination")
    checkIn := c.GetString("t-start")
    checkOut := c.GetString("t-end")

    url := "https://booking-com13.p.rapidapi.com/stays/properties/list-v2" +
        "?location=" + destination +
        "&checkin_date=" + checkIn +
        "&checkout_date=" + checkOut +
        "&language_code=en-us&currency_code=USD"

    // Create an HTTP request
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        c.Data["Error"] = "Error creating request"
        return
    }

    // Set headers required for the RapidAPI endpoint
    req.Header.Add("X-RapidAPI-Key", "3ad041d245msh9e23b1e8792896cp1fbc15jsn5e20851b7ca7")
    req.Header.Add("X-RapidAPI-Host", "booking-com13.p.rapidapi.com")

    dataChannel := make(chan models.HotelInfo)

    // Launch a goroutine to make the HTTP request and parse the response
    go func() {
        // Make the HTTP request
        res, err := http.DefaultClient.Do(req)
        if err != nil {
            c.Data["Error"] = "Error making the request"
            dataChannel <- models.HotelInfo{}
            return
        }
        defer res.Body.Close()

        // Check the HTTP status code
        if res.StatusCode != http.StatusOK {
            c.Data["Error"] = "API request failed with status code " + fmt.Sprint(res.StatusCode)
            dataChannel <- models.HotelInfo{}
            return
        }

        // Read the response body
        body, err := io.ReadAll(res.Body)
        if err != nil {
            c.Data["Error"] = "Error reading the response"
            dataChannel <- models.HotelInfo{}
            return
        }

        // Parse the JSON response into a struct
        var hotelResponse models.HotelInfo
        if err := json.Unmarshal(body, &hotelResponse); err != nil {
            c.Data["Error"] = "Error parsing JSON response"
            dataChannel <- models.HotelInfo{}
            return
        }

        // Send the extracted data to the channel
        dataChannel <- hotelResponse
    }()

    // Wait for the goroutine to complete and receive the data
    extractedData := <-dataChannel

    if extractedData.Data == nil {
        // Handle the error here
        c.Data["Error"] = "Failed to fetch hotel data"
        c.TplName="error/notFound.tpl"
        return
    }

    // Extract idDetails
    var idDetails,imgUrls, Titles, Prices []string
    for _, info := range extractedData.Data {
        idDetails = append(idDetails, info.IDDetail)
		imgUrls = append(imgUrls, info.BasicPropertyData.Photos.Main.Urls.ImgURL)
		Titles = append(Titles, info.DisplayName.Title)
		Prices = append(Prices, info.PriceDisplayInfoIrene.DisplayPrice.AmountPerStay.AmountRounded)
    }
    c.Data["Hotel"] = extractedData
	
    c.TplName = "hotel/hotelList.tpl"
}


