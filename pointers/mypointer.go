package main

import "strings"

/*
INTRODUCTION TO POINTERS
As we have learned, a variable is a named location in memory that stores a value. We can manipulate the value of a variable by assigning a new value to it or by performing operations on it. When we assign a value to a variable, we are storing that value in a specific location in memory.

x := 42
// "x" is the name of a location in memory. That location is storing the integer value of 42

A POINTER IS A VARIABLE
A pointer is a variable that stores the memory address of another variable. This means that a pointer "points to" the location of where the data is stored NOT the actual data itself.

The * syntax defines a pointer:

var p *int

The & operator generates a pointer to its operand.

myString := "hello"
myStringPtr := &myString


Unlike C, Go has no pointer arithmetic
*/

/*

Assignment


Complete the removeProfanity function.

It should use the strings.ReplaceAll function to replace all instances of the following words in the input message with asterisks.

"fubb" -> "****"
"shiz" -> "****"
"witch" -> "*****"
It should mutate the value in the pointer and return nothing.

Do not alter the function signature
*/

func removeProfanity(message *string) {
	messageVal := *message

	if message == nil {
		return
	}

	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
}

/*
POINTER RECEIVERS
A receiver type on a method can be a pointer.

Methods with pointer receivers can modify the value to which the receiver points. Since methods often need to modify their receiver, pointer receivers are more common than value receivers. However, methods with pointer receivers don't require that a pointer is used to call the method. The pointer will automatically be derived from the value.

POINTER RECEIVER
type car struct {
	color string
}

func (c *car) setColor(color string) {
	c.color = color
}

func main() {
	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)
	// prints "blue"
}
Copy icon
NON-POINTER RECEIVER
type car struct {
	color string
}

func (c car) setColor(color string) {
	c.color = color
}

func main() {
	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)
	// prints "white"
}

*/
