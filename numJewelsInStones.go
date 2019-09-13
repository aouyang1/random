package random

func numJewelsInStones(J, S string) int {
	var cnt int

	jewels := make(map[rune]struct{})

	for _, jewel := range J {
		if _, exists := jewels[jewel]; !exists {
			jewels[jewel] = struct{}{}
		}
	}

	for _, stone := range S {
		if _, exists := jewels[stone]; exists {
			cnt++
		}
	}
	return cnt
}
