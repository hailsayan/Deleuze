package main

import "fmt"

const (
	monthDay  int = 31
	monthDay2 int = 30
	monthDay3 int = 29
)

func main() {
	var month int
	var totalDays int = 0

	fmt.Println("eneter a month: ")
	fmt.Scanln(&month)

	switch month {
	case 12:
		totalDays += monthDay3
		fallthrough
	case 11:
		totalDays += monthDay2
		fallthrough
	case 10:
		totalDays += monthDay2
		fallthrough
	case 9:
		totalDays += monthDay2
		fallthrough
	case 8:
		totalDays += monthDay2
		fallthrough
	case 7:
		totalDays += monthDay2
		fallthrough
	case 6:
		totalDays += monthDay
		fallthrough
	case 5:
		totalDays += monthDay
		fallthrough
	case 4:
		totalDays += monthDay
		fallthrough
	case 3:
		totalDays += monthDay
		fallthrough
	case 2:
		totalDays += monthDay
		fallthrough
	case 1:
		totalDays += monthDay
	}

	println("total days: ", totalDays)
}
