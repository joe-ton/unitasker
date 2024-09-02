package main

import (
	"fmt"
	"time"
)

// Function to get current time and add a set number of minutes to it
func getCurrentTimeWithAddedMinutes(minutesToAdd int) string {
	currentTime := time.Now()

	// Add the specified number of minutes to the current time
	newTime := currentTime.Add(time.Duration(minutesToAdd) * time.Minute)

	zone, _ := newTime.Zone() // Get the time zone name, ignore the offset

	formattedTime := newTime.Format("03:04:05 PM " + zone)

	return formattedTime
}

func askUserForAdditionalTime() (string, error) {
	var minutesToAdd int
	fmt.Print("Add time (in minutes): ")

	_, err := fmt.Scan(&minutesToAdd)
	if err != nil {
		return "", err
	}

	return getCurrentTimeWithAddedMinutes(minutesToAdd), nil
}

func main() {
	res, err := askUserForAdditionalTime()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Println("New Time:", res)
}
