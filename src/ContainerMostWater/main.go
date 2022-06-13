package main

import "fmt"

//func maxArea(container []int) int {
//	max, start, end := 0, 0, len(container)-1
//	for start < end {
//		width := end - start
//		height := 0
//
//		if container[start] < container[end] {
//			height = container[end]
//			start++
//		} else {
//			height = container[start]
//			end--
//		}
//		temp := width * height
//
//		if temp > max {
//			max = temp
//		}
//	}
//	return max
//}

func maxArea(arr []int) int {
	max, start, end := 0, 0, len(arr)-1
	for start < end {
		width := arr[end] - arr[start]
		height := 0
		if arr[start] < arr[end] {
			height = arr[end]
			start++
		} else {
			height = arr[start]
			end--
		}
		tem := width * height
		if tem > max {
			max = tem
		}
	}
	return max
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 15, 4, 5, 9, 10}
	fmt.Println(maxArea(arr))
}
