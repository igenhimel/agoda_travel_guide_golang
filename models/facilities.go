package models

type FacilityInstance struct {
    Typename string `json:"__typename"`
    Title    string `json:"title"`
}

type FacilityData struct {
    Instances []FacilityInstance `json:"instances"`
}

type FacilityResponse struct {
    Data struct {
        Facilities []FacilityData `json:"facilities"`
    } `json:"data"`
    Message string `json:"message"`
    Status  bool   `json:"status"`
}
