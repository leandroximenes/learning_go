package expenses

import (
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	filteredItems := make([]Record, 0, len(in))
	for _, item := range in {
		matched := predicate(item)
		if matched {
			filteredItems = append(filteredItems, item)
		}
	}
	return filteredItems
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	// A function that capture the env where it was created.
	return func(r Record) bool {
		if r.Day >= p.From && r.Day <= p.To {
			return true
		}
		return false
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var total float64 = 0
	for _, item := range in {
		if item.Day >= p.From && item.Day <= p.To {
			total += item.Amount
		}

	}

	return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, category string) (float64, error) {
	categoryExpenses := Filter(in, ByCategory(category))

	if len(categoryExpenses) == 0 {
		return 0, fmt.Errorf("unknown category %s", category)
	}

	total := TotalByPeriod(categoryExpenses, p)
	return total, nil

}
