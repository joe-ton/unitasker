package main

import (
	"fmt"
	"strconv"
	"time"
)

// func getCurrentTimeWithAddedMinutes(minutesToAdd int) string {
// 	currentTime := time.Now()

// 	newTime := currentTime.Add(time.Duration(minutesToAdd) * time.Minute)

// 	zone, _ := newTime.Zone()

// 	formattedTime := newTime.Format("03:04 PM " + zone)

// 	return formattedTime
// }

// func askUserForAdditionalTime() (string, error) {
// 	var minutesToAdd int
// 	fmt.Print("Add time (in minutes): ")

// 	_, err := fmt.Scan(&minutesToAdd)
// 	if err != nil {
// 		return "", err
// 	}

// 	return getCurrentTimeWithAddedMinutes(minutesToAdd), nil
// }

func askUserForMinutes() (string, error) {
	var minutesToAdd int
	fmt.Print("Add time (in minutes): ")

	_, err := fmt.Scan(&minutesToAdd)
	if err != nil {
		return "", err
	}

	// Convert the integer to a string before returning
	return strconv.Itoa(minutesToAdd), nil
}

func setCountDownTimer(minutes int) {
	// Convert minutes to duration in seconds
	countdownDuration := time.Duration(minutes) * time.Minute

	// Create a new timer with the specified countdown duration
	timer := time.NewTimer(countdownDuration)

	// Inform the user that the countdown has started
	fmt.Printf("Countdown started for %v minutes...\n", minutes)

	// Block until the timer sends a signal on its channel
	<-timer.C

	// Notify the user when the countdown has finished
	fmt.Println("Time's up!")
}

func main() {
	// Ask User for time
	res, err := askUserForMinutes()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New Time:", res)

	// Convert the string result to an integer
	minutes, err := strconv.Atoi(res)
	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		return
	}

	// Call the countdown function with the converted minutes
	setCountDownTimer(minutes)
}
