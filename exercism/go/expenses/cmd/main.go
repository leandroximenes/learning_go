package main

import (
	"expenses"
	"fmt"
)

func Day1Expenes(r expenses.Record) bool {
	isDay1 := r.Day == 1
	return isDay1
}

var testExpensesRecords = []expenses.Record{
	{
		Day:      1,
		Amount:   5.15,
		Category: "groceries",
	},
	{
		Day:      1,
		Amount:   3.45,
		Category: "groceries",
	},
	{
		Day:      13,
		Amount:   55.67,
		Category: "utility-bills",
	},
	{
		Day:      15,
		Amount:   11,
		Category: "groceries",
	},
	{
		Day:      18,
		Amount:   244.33,
		Category: "utility-bills",
	},
	{
		Day:      20,
		Amount:   300,
		Category: "university",
	},
	{
		Day:      23,
		Amount:   20.0,
		Category: "groceries",
	},
	{
		Day:      25,
		Amount:   24.65,
		Category: "groceries",
	},
	{
		Day:      30,
		Amount:   1300,
		Category: "rent",
	},
}

func main() {
	fmt.Println("Let's test")
	period1 := expenses.DaysPeriod{From: 1, To: 15}
	period2 := expenses.DaysPeriod{From: 25, To: 26}
	period3 := expenses.DaysPeriod{From: 40, To: 50}

	records := []expenses.Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}

	filter1 := expenses.Filter(records, Day1Expenes)
	fmt.Println(filter1)

	filter2 := expenses.Filter(records, expenses.ByDaysPeriod(period3))
	fmt.Println(filter2)

	filter3 := expenses.Filter(records, expenses.ByCategory("rent"))
	fmt.Println(filter3)

	filter4 := expenses.TotalByPeriod(testExpensesRecords, period2)
	fmt.Println(filter4)

	filter5, err := expenses.CategoryExpenses(testExpensesRecords, period1, "groceries")
	fmt.Println(filter5, err)

}
