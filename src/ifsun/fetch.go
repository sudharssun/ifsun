package main
import (
    "encoding/json"
    "net/http"
    "io/ioutil"
    "errors"
)

type IfSunResponse struct {
    Parameters IfSunParameters `json:"results"`
    Status string `json:"status"`
}

type IfSunParameters struct {
    Sunrise string `json:"sunrise"`
    Sunset string  `json:"sunset"`
}

type IpVigilanteResponse struct {
    Status string `json:"status"`
    IpParameters IpParameters `json:"data"`    
}

type IpParameters struct {
    Latitude string `json:"latitude"`
    Longitude string  `json:"longitude"`
    City string `json:"city_name"`
}

func DoQuery(date string, latitude string, longitude string) (IfSunResponse, error) {
    SunResponse := IfSunResponse{}
    response, err := http.Get(SUNSETAPI_BASE + LAT_KEYWORD + latitude + LONG_KEYWORD + longitude + DATE_KEYWORD + date)
    if err != nil {        
        return SunResponse, errors.New("[Get Sunset] HTTP request failed: " + err.Error())
    }
    extractedData, _ := ioutil.ReadAll(response.Body)
    unmarshallErr := json.Unmarshal([]byte(extractedData), &SunResponse)
    if unmarshallErr != nil {
        return SunResponse, errors.New("[Get Sunset] Error occurred while parsing the response: " + unmarshallErr.Error())
    }

    return SunResponse, nil
}

func GetSunset(date string, lat string, long string) (string, error) {
    SunResponse, err := DoQuery(date, lat, long)
    if err != nil {
        return ERROR, err
    }

    return SunResponse.Parameters.Sunset, nil
}

func GetCurrentLocation() (string, string, error) {
    IpResponse := IpVigilanteResponse{}
    response, err := http.Get(IPVIGILANTEAPI_BASE)
    if(err != nil) {
        return EMPTY_STRING, EMPTY_STRING, errors.New("[Get Sunset] Failed to get location. Try providing a city name as an argument: " + err.Error())
    }

    extractedData, _ := ioutil.ReadAll(response.Body)
    unmarshallErr := json.Unmarshal([]byte(extractedData), &IpResponse)
    if unmarshallErr != nil {        
        return EMPTY_STRING, EMPTY_STRING, errors.New("[Get Sunset] Failed to parse response to get location:" + unmarshallErr.Error())
    }

    return IpResponse.IpParameters.Latitude, IpResponse.IpParameters.Longitude, nil;
}