package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/cheshire137/gogostardewcal/pkg/calendar"
)

func main() {
	season, err := getSeason()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	day, err := getDay(season)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	calendar, err := calendar.NewCalendar("pkg/calendar/calendar.json", day, season)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(calendar)

	events, err := calendar.CurrentEvents()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	totalEvents := len(events)
	if totalEvents < 1 {
		fmt.Println("No events today")
		os.Exit(0)
	}

	plural := "s"
	if totalEvents == 1 {
		plural = ""
	}
	fmt.Printf("%d event%s today:\n", totalEvents, plural)
	for _, event := range events {
		fmt.Printf("- %s\n", event)
	}

	fmt.Println(calendar.DaySheet())
}

var seasons = []string{
	"spring",
	"summer",
	"fall",
	"winter",
}

func getSeason() (string, error) {
	fmt.Println("What is the current season in Stardew Valley?")
	fmt.Println("1) 🌸 spring")
	fmt.Println("2) 🌻 summer")
	fmt.Println("3) 🍄 fall")
	fmt.Println("4) ⛄️ winter")
	fmt.Println("5) Exit")
	fmt.Printf("> ")
	var seasonChoice int
	_, err := fmt.Scanf("%d", &seasonChoice)
	if err != nil {
		return "", err
	}
	if seasonChoice == 5 {
		fmt.Println("Exiting...")
		os.Exit(0)
	}
	if seasonChoice < 1 || seasonChoice > 4 {
		return "", errors.New("Error: invalid choice, please choose between 1-4")
	}
	return seasons[seasonChoice-1], nil
}

func getDay(season string) (int, error) {
	fmt.Printf("What day of %s is it?\n", season)
	fmt.Println("Enter 1-28:")
	fmt.Printf("> ")
	var dayChoice int
	_, err := fmt.Scanf("%d", &dayChoice)
	if err != nil {
		return 0, err
	}
	if dayChoice < 1 || dayChoice > 28 {
		return 0, errors.New("Error: invalid choice, please choose between 1-28")
	}
	return dayChoice, nil
}
