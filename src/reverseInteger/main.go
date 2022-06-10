package main

import "fmt"

// 123 %10  = 3   0*10 + 3 = 3
// 12 %10 = 2 	  3*10 + 2 = 32
// 1 % 10 = 1     32*10 +1 = 321
// 将数调转
func reverse7(x int) int {
	tmp := 0
	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10
		fmt.Println(tmp)
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}

func main() {
	var x int
	x = 1234567
	fmt.Println(reverse7(x))
}
