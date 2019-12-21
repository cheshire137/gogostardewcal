package main

import (
	"fmt"
	"os"
)

var seasons = []string{
	"spring",
	"summer",
	"fall",
	"winter",
}

func main() {
	fmt.Println("What is the current season in Stardew Valley?")
	fmt.Println("1) Spring")
	fmt.Println("2) Summer")
	fmt.Println("3) Fall")
	fmt.Println("4) Winter")
	var seasonChoice int
	_, err := fmt.Scanf("%d", &seasonChoice)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if seasonChoice < 1 || seasonChoice > 4 {
		fmt.Println("Error: invalid choice, please choose between 1-4")
		os.Exit(1)
	}
	season := seasons[seasonChoice-1]
	fmt.Println(season)
}
