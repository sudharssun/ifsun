package main

import (
    "net/http"
    "os"
    "io/ioutil"
    "net/url"
    "fmt"
    "errors"
    "encoding/json"
)

type GCResponse struct {
    OuterResponse []GCOuterResponse `json:"results"`
}

type GCOuterResponse struct {
    Results GCAnnotations `json:"annotations"`
    Geometry Geometry `json:"geometry"`
}

type GCAnnotations struct {    
    TimeZone TimeZone `json:"timezone"`    
}

type TimeZone struct {
    Name string `json:"name"`
}

type Geometry struct {
    Latitude  float64 `json:"lat"`
    Longitude float64 `json:"lng"`
}

func GetCoordinates(args []string) (string, string, string, error) {    
    GcResponse := GCResponse{}
    
    u, _ := url.Parse(OPENCAGEAPI_BASE)
    u.Path += "json"

    q := u.Query()
    q.Set("q", args[0])
    q.Set("key", os.Getenv(API_KEY))
    
    u.RawQuery = q.Encode()

    response, err := http.Get(u.String())
    if(err != nil) {
        return EMPTY_STRING, EMPTY_STRING, EMPTY_STRING, errors.New("[Get coordinates] Http request failed: " + err.Error())
    }

    extractedData, _ := ioutil.ReadAll(response.Body)
    unmarshallErr := json.Unmarshal([]byte(extractedData), &GcResponse)

    if unmarshallErr != nil {
        return EMPTY_STRING, EMPTY_STRING, EMPTY_STRING, errors.New("[Get coordinates] Error occurred while parsing the response for getting coordinates: " + unmarshallErr.Error())
    }

    if len(GcResponse.OuterResponse) < 1 {
        return EMPTY_STRING, EMPTY_STRING, EMPTY_STRING, errors.New("[Get coordinates] Place Not found. Misspelled? Please Try again")
    }
    
    // The API can respond with multiple matches with the top most one as the most relevant to the search query.
    return fmt.Sprintf("%f", GcResponse.OuterResponse[0].Geometry.Latitude), fmt.Sprintf("%f", GcResponse.OuterResponse[0].Geometry.Longitude), string(GcResponse.OuterResponse[0].Results.TimeZone.Name), nil
}