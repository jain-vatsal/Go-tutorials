package main

import "fmt"

func bulkSend(numMessages int) float64 {
	// ?
	cost := 0.0
	vr := 0.00

	for i := 0; i < numMessages; i++ {
		cost = cost + vr
		vr = vr + 0.01
	}

	return cost
}

func maxMessages(thresh int) int {
	cost := 0
	ans := 0
	for i := 0; ; i++ {
		if cost+i*1 <= thresh {
			ans++
			cost = cost + 1*i
		} else {
			break
		}
	}
	return ans
}

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Println("fizz")
		} else if i%3 != 0 && i%5 == 0 {
			fmt.Println("buzz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else {
			fmt.Println(i)
		}
	}
}

func printPrimes(max int) {
	for n := 2; n <= max; n++ {
		if n == 2 {
			fmt.Println(n)
		}
		if n%2 == 0 {
			continue
		}
		isPrime := true
		for i := 3; i*i <= n; i++ {
			if n%i == 0 {
				isPrime = false
				continue
			}
		}

		if !isPrime {
			continue
		}
		fmt.Println(n)

	}
}
