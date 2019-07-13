package main
import (
    "fmt"
    "encoding/json"
    "net/http"
    "io/ioutil"
    "time"
    "github.com/araddon/dateparse"
)

type IfSunResponse struct {
    Parameters Parameters `json:"results"`
    Status string `json:"status"`
}

type Parameters struct {
    Sunrise string `json:"sunrise"`
    Sunset string  `json:"sunset"`
    SolarNoon string `json:"solar_noon"`
    DayLength string `json:"day_length"`
    CivilTwlBegin string `json:"civil_twilight_begin"`
    CivilTwlEnd string `json:"civil_twilight_end"`
    NauticalTwlBegin string `json:"nautical_twilight_begin"`
    NauticalTwlEnd string `json:"nautical_twilight_end"`
    AstTwlBegin string `json:"astronomical_twilight_begin"`
    AstTwlEnd string `json:"astronomical_twilight_end"`
}

func DoQuery(date string, latitude string, longitude string) (IfSunResponse, error) {
    SunResponse := IfSunResponse{}
    response, err := http.Get("https://api.sunrise-sunset.org/json?lat=" + latitude + "&lng=" + longitude + "&date=" + date)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
        return SunResponse, err
    }
    extractedData, _ := ioutil.ReadAll(response.Body)
    unmarshallErr := json.Unmarshal([]byte(extractedData), &SunResponse)
    if unmarshallErr != nil {
        fmt.Printf("Error occurred while parsing the response %s\n", unmarshallErr)
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

func main() {
    // Get sunset time in UTC
    loc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err.Error())
	}
	time.Local = loc
    
    // Hardcoded to Seattle's Lat and Long - TODO make city as arg and convert city to lat and long
    SunsetToday, _ := GetSunset(GetDate(0), "47.6038321", "-122.3300624")
    SunsetTom, _ := GetSunset(GetDate(1), "47.6038321", "-122.3300624")    
    SunsetDayAfter, _ := GetSunset(GetDate(2), "47.6038321", "-122.3300624")  

    TimeFormatToday, errToday := dateparse.ParseLocal(GetDate(0) + " " + SunsetToday)    
    TimeFormatTom, errTom := dateparse.ParseLocal(GetDate(1) + " " + SunsetTom)
    TimeFormatDayAfter, errDayAfter := dateparse.ParseLocal(GetDate(2) + " " + SunsetDayAfter)

    if errToday != nil || errTom != nil || errDayAfter != nil {
        fmt.Printf("Err1: %s, err2: %s\n", errToday.Error(), errTom.Error())
        return
    }

    // Print time in local timezone
    localLoc, err:=time.LoadLocation("America/Los_Angeles")
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Today: " + TimeFormatToday.In(localLoc).Format("15:04"))
    fmt.Println("Tomorrow: " + TimeFormatTom.In(localLoc).Format("15:04"))
    fmt.Println("Day after: " + TimeFormatDayAfter.In(localLoc).Format("15:04"))

    fmt.Println("(Credit: https://sunrise-sunset.org/api)")
}
