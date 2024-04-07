/*
THERE IS NO WHILE LOOP IN GO
Most programming languages have a concept of a while loop. Because Go allows for the omission of sections of a for loop, a while loop is just a for loop that only has a CONDITION.

for CONDITION {
  // do some stuff while CONDITION is true
}

For example:

plantHeight := 1
for plantHeight < 5 {
  fmt.Println("still growing! current height:", plantHeight)
  plantHeight++
}
fmt.Println("plant has grown to ", plantHeight, "inches")
*/

/*
We have an interesting new cost structure from our SMS vendor. They charge exponentially more money for each consecutive text we send! Let's write a function that calculates how many messages we can send in a given batch given a costMultiplier and a maxCostInPennies.

In a nutshell, the first message costs a penny, and each message after that costs the same as the previous message multiplied by the costMultiplier.

There is a bug in the code! Add a condition to the for loop to fix the bug. The loop should stop when actualCostInPennies is greater than or equal to the maxCostInPennies.
*/

package main

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 0
	for actualCostInPennies <= float64(maxCostInPennies) {
		actualCostInPennies *= costMultiplier
		maxMessagesToSend++
	}
	if actualCostInPennies > float64(maxCostInPennies) {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}
