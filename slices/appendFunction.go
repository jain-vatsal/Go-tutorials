package main

/*
APPEND
The built-in append function is used to dynamically add elements to a slice:

func append(slice []Type, elems ...Type) []Type

If the underlying array is not large enough, append() will create a new underlying array and point the slice to it.

Notice that append() is variadic, the following are all valid:

slice = append(slice, oneThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice...)
*/

/*

ASSIGNMENT
We've been asked to "bucket" costs per day, in a given period.

Complete the getCostsByDay() function using the append() function. It accepts a slice of cost structs and returns a slice of float64s, where each float64 represents the total cost for a day.

The length of the daily costs slice should be the number of days contained in the costs slice, up to and including the last day. There should only be one "bucket" of costs per day. Be sure to include entries for intermediate days, even if they don't have any costs.

Days are numbered and start at 0.

EXAMPLE
Given this input:

[]cost{
    {0, 4.0},
    {1, 2.1},
    {5, 2.5},
    {1, 3.1},
}

We expect this result:

[]float64{
    4.0, // first day
    5.2, // 2.1 + 3.1
    0.0, // intermediate days with no costs
    0.0, // ...
    0.0, // ...
    2.5, // last day
}*/

type cost struct {
	day   int
	value float64
}

func returnDays(costs []cost) int {
	ans := 0
	for i := 0; i < len(costs); i++ {
		if costs[i].day > ans {
			ans = costs[i].day
		}
	}
	return ans
}

func getCostsByDay(costs []cost) []float64 {
	// ?

	days := returnDays(costs)
	costsPerDay := make([]float64, days+1)

	for i := 0; i < len(costsPerDay); i++ {
		costsPerDay[i] = 0.0
	}

	for i := 0; i < len(costs); i++ {
		idx := costs[i].day
		costsPerDay[idx] += costs[i].value
	}

	return costsPerDay
	return nil
}
