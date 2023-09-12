package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"travel_guide/models"

	beego "github.com/beego/beego/v2/server/web"
)





type HotelControllerDetails struct {
    beego.Controller
}
func (c *HotelControllerDetails) Get() {
    idDetail := c.GetString("IDDetail")
    cleanedIdDetail := strings.ReplaceAll(idDetail, "es/", "")
    cleanedIdDetail = strings.ReplaceAll(cleanedIdDetail, "-", " ")

    // Capitalize the first letter of each word
    words := strings.Fields(cleanedIdDetail)
    for i, word := range words {
        words[i] = strings.Title(word)
    }
    cleanedIdDetail = strings.Join(words, " ")
    languageCode := "en-us"

    // Get Photo URIs
    photoURIs, err := GetPhotoURIs(idDetail, languageCode)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Get Description
    description, err := GetDescription(idDetail, languageCode)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Get Facility Titles
    titles, err := GetFacilityTitles(idDetail, languageCode)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Set data in the c.Data map
    c.Data["PhotoURIs"] = photoURIs
    c.Data["Description"] = description
    c.Data["FacilityTitles"] = titles
    c.Data["HotelTitle"] = cleanedIdDetail

    

    // Set the template name
    c.TplName = "hotel/hotel_detail.tpl"
}




func GetDescription(idDetail, languageCode string) (string, error) {
    url := fmt.Sprintf("https://booking-com13.p.rapidapi.com/stays/properties/detail/description?id_detail=%s&language_code=%s", idDetail, languageCode)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return "", err
    }

    req.Header.Add("X-RapidAPI-Key", "3ad041d245msh9e23b1e8792896cp1fbc15jsn5e20851b7ca7")
    req.Header.Add("X-RapidAPI-Host", "booking-com13.p.rapidapi.com")

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return "", err
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        return "", fmt.Errorf("API request failed with status code: %d", res.StatusCode)
    }

    // Read the response body
    body, err := io.ReadAll(res.Body)
    if err != nil {
        return "", err
    }

    // Unmarshal the JSON response into a Response struct
    var response models.Response
    if err := json.Unmarshal(body, &response); err != nil {
        return "", err
    }

    // Check if there is at least one description in the response
    if len(response.Data) > 0 {
        return response.Data[0].Description, nil
    }

    return "No description found in the response.", nil
}


// GetPhotoURIs retrieves photo URIs for a given idDetail
func GetPhotoURIs(idDetail, languageCode string) ([]string, error) {
    url := fmt.Sprintf("https://booking-com13.p.rapidapi.com/stays/properties/detail/photos?id_detail=%s&language_code=%s", idDetail, languageCode)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Add("X-RapidAPI-Key", "3ad041d245msh9e23b1e8792896cp1fbc15jsn5e20851b7ca7")
    req.Header.Add("X-RapidAPI-Host", "booking-com13.p.rapidapi.com")

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API request failed with status code: %d", res.StatusCode)
    }

    // Read the response body
    body, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, err
    }

    // Unmarshal the JSON response into a Response struct
    var response models.ResponsePhoto
    if err := json.Unmarshal(body, &response); err != nil {
        return nil, err
    }

    // Extract all photoUri values
    photoUris := []string{}
    for _, photo := range response.Data.Photos {
    photoUris = append(photoUris, photo.PhotoUri)
    if len(photoUris) >= 4 {
        break // Stop after collecting 4 URIs
    }
}

    return photoUris, nil
}

func GetFacilityTitles(idDetail, languageCode string) ([]string, error) {
    url := fmt.Sprintf("https://booking-com13.p.rapidapi.com/stays/properties/detail/facilities?id_detail=%s&language_code=%s", idDetail, languageCode)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Add("X-RapidAPI-Key", "3ad041d245msh9e23b1e8792896cp1fbc15jsn5e20851b7ca7")
    req.Header.Add("X-RapidAPI-Host", "booking-com13.p.rapidapi.com")

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API request failed with status code: %d", res.StatusCode)
    }

    // Read the response body
    body, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, err
    }

    // Unmarshal the JSON response into a FacilityResponse struct
    var response models.FacilityResponse
    if err := json.Unmarshal(body, &response); err != nil {
        return nil, err
    }

    // Extract at least 5 titles from the response
    titles := []string{}
    for _, facility := range response.Data.Facilities {
        for _, instance := range facility.Instances {
            if len(titles) >= 5 {
                return titles, nil
            }
            titles = append(titles, instance.Title)
        }
    }

    return titles, nil
}
