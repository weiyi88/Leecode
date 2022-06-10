package main

import "fmt"

// 滑动窗口
// ASCII码，127位
// 只要没有重复，freq数组该ASCII码的位值+1
// 遇到重复，freq数组该位有值，则重复
// 这时候，right +1 就是最大值
// 依次遍历字符串，直到字符串结束
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	result, left, right := 0, 0, -1
	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var str = "abcab"
	fmt.Println(lengthOfLongestSubstring(str))
}
