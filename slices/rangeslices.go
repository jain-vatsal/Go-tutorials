package main

import "fmt"

// fruits := []string{"apple", "banana", "grape"}
// for i, fruit := range fruits {
//     fmt.Println(i, fruit)
// }

func indexOfFirstBadWord(msg []string, badWords []string) int {

	// for i := 0; i < len(msg); i++ {
	// 	for j := 0; j < len(badWords); j++ {
	// 		if msg[i] == badWords[j] {
	// 			return i
	// 		}
	// 	}
	// }

	for i, word := range msg {
		for j, badWord := range badWords {
			if word == badWord {
				return i
			}
			fmt.Println(j)
		}
	}
	return -1
}
