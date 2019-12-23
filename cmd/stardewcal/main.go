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

	events, err := calendar.CurrentEvents()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines := calendar.EventsSummary(events)
	fmt.Println(calendar.DaySheet(lines...))
}

var seasons = []string{
	"spring",
	"summer",
	"fall",
	"winter",
}

func getSeason() (string, error) {
	fmt.Println("What is the current season in Stardew Valley?")
	fmt.Printf("1) %s spring\n", calendar.SPRING_EMOJI)
	fmt.Printf("2) %s summer\n", calendar.SUMMER_EMOJI)
	fmt.Printf("3) %s fall\n", calendar.FALL_EMOJI)
	fmt.Printf("4) %s winter\n", calendar.WINTER_EMOJI)
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
