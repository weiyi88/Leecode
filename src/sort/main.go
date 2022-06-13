package main

import "fmt"

func Bubble(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{22, 34, 3, 32, 82, 55, 89, 50, 37, 5, 64, 35, 9, 70}
	fmt.Println(Bubble(arr))
}
