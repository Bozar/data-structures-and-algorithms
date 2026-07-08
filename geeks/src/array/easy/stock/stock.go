// https://www.geeksforgeeks.org/dsa/stock-buy-sell/
package stock

import ()

func Print(input []int) int {
	if len(input) < 2 {
		return 0
	}

	profit := 0
	idxBuy := -1
	for i := range input {
		if i == 0 {
			continue
		}
		if input[i] >= input[i-1] {
			if idxBuy < 0 {
				idxBuy = i - 1
			}
			continue
		}
		if idxBuy < 0 {
			continue
		}
		profit += input[i-1] - input[idxBuy]
		idxBuy = -1
	}
	if idxBuy >= 0 {
		profit += input[len(input)-1] - input[idxBuy]
	}
	return profit
}
