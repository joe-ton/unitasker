package main

import (
	"fmt"
	"strings"
	"time"
)

func getCurrentTime() string {
	currentTime := time.Now()
	zone, _ := currentTime.Zone() // Get the time zone name, ignore the offset

	formattedTime := currentTime.Format("03:04:05 PM " + zone)

	// Replace "PDT" with "PST" if "PDT" is found in the formatted time string
	updatedTimeString := strings.Replace(formattedTime, "PDT", "PST", 1)

	return updatedTimeString
}

func main() {
	res := getCurrentTime()
	fmt.Println(res)
}
