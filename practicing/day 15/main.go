package main

import "fmt"

type Record struct {
	Day int
}

type DaysPeriodFilter struct {
	From int
	To   int
}

type Filter interface {
	Match(Record) bool
}

func (f DaysPeriodFilter) Match(r Record) bool {
	return r.Day >= f.From && r.Day <= f.To
}

func ApplyFilter(records []Record, f Filter) []Record {
	result := make([]Record, 0, len(records))
	for _, item := range records {
		matched := f.Match(item)
		if matched {
			result = append(result, item)
		}
	}

	return result
}

func main() {
	fmt.Println("Reviewing")

	period := DaysPeriodFilter{10, 20}
	records := []Record{
		{Day: 5},
		{Day: 15},
		{Day: 25},
	}

	result := ApplyFilter(records, period)

	fmt.Println(result)

	fmt.Println("Today's review:")
	fmt.Println("Understanding how Go uses interfaces in practical way (without useless theory)")
}
