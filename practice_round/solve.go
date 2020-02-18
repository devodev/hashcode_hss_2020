package main

// Reverse .
func Reverse(s []int) {
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
}

// Solve .
func Solve(m, n int, payload []int) []int {
	result := []int{}
	for i := n - 1; i >= 0; i-- {
		subResult := []int{}
		val := payload[i]
		if val <= m {
			subResult = append(subResult, i)
			if i > 0 {
				solved := Solve(m-val, i, payload[0:i])
				subResult = append(subResult, solved...)
			}
			resultSum := sum(result, payload)
			subResultSum := sum(subResult, payload)
			if subResultSum > resultSum {
				result = subResult
			}
			// Fast exit
			if subResultSum == m {
				break
			}
		}
	}
	return result
}

func sum(indexes, payload []int) int {
	result := 0
	for _, v := range indexes {
		result += payload[v]
	}
	return result
}
