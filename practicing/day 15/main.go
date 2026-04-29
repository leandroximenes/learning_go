package main

import "fmt"

type Record struct {
	Day int
}

type DaysPeriod struct {
	From int
	To   int
}

func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		result := r.Day >= p.From && r.Day <= p.To
		return result
	}
}

func Filter(records []Record, fn func(Record) bool) []Record {
	result := make([]Record, 0, len(records))
	for _, item := range records {
		matched := fn(item)
		if matched {
			result = append(result, item)
		}
	}

	return result
}

func main() {
	fmt.Println("Reviewing")

	period := DaysPeriod{10, 20}
	records := []Record{
		{Day: 5},
		{Day: 15},
		{Day: 25},
	}

	for _, item := range records {
		fmt.Printf("%p\n", &item)
	}
	
	fmt.Printf("Array: %p\n", &records)
	
	result := Filter(records, ByDaysPeriod(period))

	fmt.Println(result)
}
