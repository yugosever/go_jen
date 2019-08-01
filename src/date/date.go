package main

import (
	"fmt"
	"log"
	"calendar"
)

func main(){
	date := calendar.Event{}
	date.Title = "New Year"
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)
	fmt.Println(date.Year())
	fmt.Println(date.Title)
}
