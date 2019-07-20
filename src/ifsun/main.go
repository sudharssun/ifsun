package main
import (
    "time"
    "github.com/araddon/dateparse"
    "strings"
    "fmt"
    "os"
    "io/ioutil"
)

// Usage: ifsun [arg]
//        arg: city name, address or country
//        Examples: ifsun Dubai
//                  ifsun France
func main() {
    args := os.Args[1:];
    lat, long, err := GetCurrentLocation()
    exitIfError(err)

    timezoneinBytes, _ := ioutil.ReadFile(TIMEZONE_FILE)
    timezone := string(timezoneinBytes)
    timezone = strings.TrimSuffix(timezone, "\n")
    
    if(len(args) > 0){
        lat, long, timezone, err = GetCoordinates(args)
        exitIfError(err)
    }

    // Get sunset time in UTC
    loc, err := time.LoadLocation(UTC)
    exitIfError(err)

    time.Local = loc
    SunsetToday, _ := GetSunset(CurrentDateWithOffset(0), lat, long)
    SunsetTom, _ := GetSunset(CurrentDateWithOffset(1), lat, long)    
    SunsetDayAfter, _ := GetSunset(CurrentDateWithOffset(2), lat, long)  

    TimeFormatToday, errToday := dateparse.ParseLocal(CurrentDateWithOffset(0) + " " + SunsetToday)    
    TimeFormatTom, errTom := dateparse.ParseLocal(CurrentDateWithOffset(1) + " " + SunsetTom)
    TimeFormatDayAfter, errDayAfter := dateparse.ParseLocal(CurrentDateWithOffset(2) + " " + SunsetDayAfter)
    
    exitIfError(errToday)
    exitIfError(errTom)
    exitIfError(errDayAfter)

    // Print sunset time in timezone of the location
    localLoc, err:=time.LoadLocation(timezone)
    exitIfError(err)
    
    fmt.Println("Today: " + TimeFormatToday.In(localLoc).Format("15:04"))
    fmt.Println("Tomorrow: " + TimeFormatTom.In(localLoc).Format("15:04"))
    fmt.Println("Day after: " + TimeFormatDayAfter.In(localLoc).Format("15:04"))
}
