package main

func findSum(x, y int) int {
	return x + y
}

func findProd(x, y int) int {
	return x * y
}

func aggregate(a, b, c int, arith func(int, int) int) int {
	return arith(arith(a, b), c)
}

// func main() {
// 	ans := aggregate(1, 2, 4, findSum)
// 	fmt.Println(ans)

// 	ans = aggregate(1, 2, 4, findProd)
// 	fmt.Println(ans)
// }
