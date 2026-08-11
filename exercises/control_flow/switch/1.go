package main

import "fmt"

func DayType(day string) string {
	// TODO: use switch to return "Weekend" for Sat/Sun, "Weekday" otherwise
	return ""
}

func main() {
	fmt.Println(DayType("Sat"))
	fmt.Println(DayType("Mon"))
}