// https://www.geeksforgeeks.org/dsa/generating-subarrays-using-recursion/
package genSub

import ()

func PrintIter(input []int) [][]int {
	var output [][]int
	for i := 0; i < len(input); i++ {
		for j := 0; j < i+1; j++ {
			output = append(output, input[j:i+1])
		}
	}
	return output
}

func PrintRec(input []int) [][]int {
	if len(input) < 2 {
		return [][]int{input}
	}
	output := subHead(input)
	output = append(output, PrintRec(input[1:])...)
	return output
}

func subHead(input []int) [][]int {
	if len(input) < 1 {
		return [][]int{}
	}
	output := [][]int{input}
	output = append(output, subHead(input[0:len(input)-1])...)
	return output
}
