// https://www.geeksforgeeks.org/dsa/the-celebrity-problem/
package celebrity

import ()

func Print(input [][]int) int {
	mSize := len(input)
	people := make([]int, mSize)
	for i := 0; i < mSize; i++ {
		for j := 0; j < mSize; j++ {
			if i == j {
				continue
			}
			if input[i][j] == 1 {
				people[i] = -1
				if people[j] > -1 {
					people[j] += 1
				}
			}
		}
	}
	for i, v := range people {
		if v == mSize-1 {
			return i
		}
	}
	return -1
}
