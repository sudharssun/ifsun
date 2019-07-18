package main
import (
    "fmt"
    "os"
    "encoding/json"
    "net/http"
    "io/ioutil"
    "time"
    "github.com/araddon/dateparse"
    "strings"
    "net/url"
)

type IfSunResponse struct {
    Parameters Parameters `json:"results"`
    Status string `json:"status"`
}

type Parameters struct {
    Sunrise string `json:"sunrise"`
    Sunset string  `json:"sunset"`
    // SolarNoon string `json:"solar_noon"`
    // DayLength string `json:"day_length"`
    // CivilTwlBegin string `json:"civil_twilight_begin"`
    // CivilTwlEnd string `json:"civil_twilight_end"`
    // NauticalTwlBegin string `json:"nautical_twilight_begin"`
    // NauticalTwlEnd string `json:"nautical_twilight_end"`
    // AstTwlBegin string `json:"astronomical_twilight_begin"`
    // AstTwlEnd string `json:"astronomical_twilight_end"`
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
    response, err := http.Get("https://api.sunrise-sunset.org/json?lat=" + latitude + "&lng=" + longitude + "&date=" + date)
    if err != nil {
        fmt.Println("The HTTP request failed with error %s\n", err)
        return SunResponse, err
    }
    extractedData, _ := ioutil.ReadAll(response.Body)
    unmarshallErr := json.Unmarshal([]byte(extractedData), &SunResponse)
    if unmarshallErr != nil {
        fmt.Println("Error occurred while parsing the response %s\n", unmarshallErr)
        return SunResponse, unmarshallErr
    }

    return SunResponse, nil
}

func GetSunset(date string, lat string, long string) (string, error) {
    SunResponse, err := DoQuery(date, lat, long)
    if err != nil {
        return "ERROR", err
    }

    return SunResponse.Parameters.Sunset, nil
}

func GetDate(offsetFromToday int) string {
    currentTime := time.Now()
    offsetTime := currentTime.AddDate(0, 0, offsetFromToday)
    return offsetTime.Format("2006-01-02")        
}

func GetCurrentLocation() (string, string) {
    IpResponse := IpVigilanteResponse{}
    response, err := http.Get("https://ipvigilante.com/")
    if(err != nil) {
        fmt.Println("Failed to get current city.Try providing a city name as an argument");
        return "", ""
    }

    extractedData, _ := ioutil.ReadAll(response.Body)
    unmarshallErr := json.Unmarshal([]byte(extractedData), &IpResponse)
    if unmarshallErr != nil {
        fmt.Println("Error occurred while parsing the response %s\n", unmarshallErr)
        return "", ""
    }
    return IpResponse.IpParameters.Latitude, IpResponse.IpParameters.Longitude;
}

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

type Geometry struct {
    Latitude  float64 `json:"lat"`
    Longitude float64 `json:"lng"`
}

type TimeZone struct {
    Name string `json:"name"`
}

func GetCoordinates(args []string) (string, string, string) {    
    GcResponse := GCResponse{}
    
    u, _ := url.Parse("https://api.opencagedata.com/geocode/v1/")
	u.Path += "json"

	q := u.Query()
	q.Set("q", args[0])
    q.Set("key", "362bbf0a11d247b88be2512d44811c0b")
    u.RawQuery = q.Encode()

    response, err := http.Get(u.String())
    if(err != nil) {
        fmt.Println("Failed to get current city.Try providing a city name as an argument");
        return "", "", ""
    }

    extractedData, _ := ioutil.ReadAll(response.Body)
    unmarshallErr := json.Unmarshal([]byte(extractedData), &GcResponse)

    if unmarshallErr != nil {
        fmt.Println("Error occurred while parsing the response %s\n", unmarshallErr)
        return "", "", ""
    }

    if len(GcResponse.OuterResponse) < 1 {
        fmt.Println("Place not found. Misspelled? Please Try again")
        return "", "", ""
    }
    
    return fmt.Sprintf("%f", GcResponse.OuterResponse[0].Geometry.Latitude), fmt.Sprintf("%f", GcResponse.OuterResponse[0].Geometry.Longitude), string(GcResponse.OuterResponse[0].Results.TimeZone.Name)
}

func main() {
    args := os.Args[1:];
    lat, long := GetCurrentLocation()
    timezoneinBytes, _ := ioutil.ReadFile("/etc/timezone")
    timezone := string(timezoneinBytes)
    timezone = strings.TrimSuffix(timezone, "\n")
    
    if(len(args) > 0){
        lat, long, timezone = GetCoordinates(args)
        if(lat == "" || long == "" || timezone == ""){
            return
        }
    }

    // Get sunset time in UTC
    loc, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println("Error loading location/timezone")
	}
	time.Local = loc
	
	
    SunsetToday, _ := GetSunset(GetDate(0), lat, long)
    SunsetTom, _ := GetSunset(GetDate(1), lat, long)    
    SunsetDayAfter, _ := GetSunset(GetDate(2), lat, long)  

    TimeFormatToday, errToday := dateparse.ParseLocal(GetDate(0) + " " + SunsetToday)    
    TimeFormatTom, errTom := dateparse.ParseLocal(GetDate(1) + " " + SunsetTom)
    TimeFormatDayAfter, errDayAfter := dateparse.ParseLocal(GetDate(2) + " " + SunsetDayAfter)

    if errToday != nil || errTom != nil || errDayAfter != nil {
        fmt.Println("Error fetching sunset time", errToday.Error(), errTom.Error())
        return
    }

    // Print time in local timezone
    localLoc, err:=time.LoadLocation(timezone)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Today: " + TimeFormatToday.In(localLoc).Format("15:04"))
    fmt.Println("Tomorrow: " + TimeFormatTom.In(localLoc).Format("15:04"))
    fmt.Println("Day after: " + TimeFormatDayAfter.In(localLoc).Format("15:04"))

    fmt.Println("(Credit: https://sunrise-sunset.org/api)")
}
