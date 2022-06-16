package basic

import (
	"fmt"
	"time"
)

func ExampleTime1() {
	utc := time.Now().UTC()
	fmt.Println(utc)
	fmt.Println(utc.Format("2006-01-02 15:04, (UTC -07:00)"))

	local := utc
	location, err := time.LoadLocation("Europe/Budapest")
	if err == nil {
		local = local.In(location)
	}
	fmt.Println("UTC", utc.Format("15:04"), local.Location(), local.Format("15:04"))
	fmt.Println(local.Format("2006-01-02 15:04, (UTC -07:00)"))

	local = utc
	location, err = time.LoadLocation("America/Los_Angeles")
	if err == nil {
		local = local.In(location)
	}
	fmt.Println("UTC", utc.Format("15:04"), local.Location(), local.Format("15:04"))
	fmt.Println(local.Format("2006-01-02 15:04, (UTC -07:00)"))

	local = utc
	location, err = time.LoadLocation("Asia/Muscat") // Arabian Standard Time
	if err == nil {
		local = local.In(location)
	}
	fmt.Println(local)
	fmt.Println("UTC", utc.Format("15:04"), local.Location(), local.Format("15:04"))
	fmt.Println(local.Format("2006-01-02 15:04, (UTC -07:00)"))

	today := time.Now()
	omanTime, omanTimeStr := GetAsiaMuscatTime(today, "2006-01-02 15:04, (UTC -07:00)")
	fmt.Println(today)
	fmt.Println(today.UTC())
	fmt.Println(omanTime)
	fmt.Println(omanTimeStr)

	// Output:
}

const ASIA_MUSCAT_LOCATION = "Asia/Muscat"

var AsiaMuscatTime *time.Location

func init() {
	location, err := time.LoadLocation(ASIA_MUSCAT_LOCATION)
	if err != nil {
		location, _ = time.LoadLocation("")
	}
	AsiaMuscatTime = location
}

func GetAsiaMuscatTime(t time.Time, layout string) (time.Time, string) {
	local := t
	local = local.In(AsiaMuscatTime)
	return local, local.Format(layout)
}

func GetFormatTime(t time.Time, loc, layout string) string {
	if loc == "" {
		loc = ASIA_MUSCAT_LOCATION
	}
	local := t
	location, err := time.LoadLocation(loc)
	if err == nil {
		local = local.In(location)
	}
	return local.Format(layout)
}
