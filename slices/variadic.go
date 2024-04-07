package main

/*
Many functions, especially those in the standard library, can take an arbitrary number of final arguments. This is accomplished by using the "..." syntax in the function signature.

A variadic function receives the variadic arguments as a slice.

func concat(strs ...string) string {
    final := ""
    // strs is just a slice of strings
    for i := 0; i < len(strs); i++ {
        final += strs[i]
    }
    return final
}

func main() {
    final := concat("Hello ", "there ", "friend!")
    fmt.Println(final)
    // Output: Hello there friend!
}
*/

/*

SPREAD OPERATOR
The spread operator allows us to pass a slice into a variadic function. The spread operator consists of three dots following the slice in the function call.

func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {
    names := []string{"bob", "sue", "alice"}
    printStrings(names...)
}
*/

func sum(nums ...int) int {
	ans := 0

	for i := 0; i < len(nums); i++ {
		ans = ans + nums[i]
	}
	return ans
}
