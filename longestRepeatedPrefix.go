package random

type node struct {
	val  byte
	cnt  int
	next nodeIndex
}

type nodeIndex map[byte]*node

func addString(str string, ni nodeIndex) {
	if len(str) == 0 {
		return
	}

	f := str[0]
	if _, exists := ni[f]; !exists {
		ni[f] = &node{f, 1, make(nodeIndex)}
	} else {
		ni[f].cnt++
	}

	if len(str) > 1 {
		addString(str[1:], ni[f].next)
	}
}

func findLongest(longestPref *string, currPref string, ni nodeIndex) {
	for _, n := range ni {
		if n.cnt < 2 {
			if longestPref != nil {
				if len(currPref) > len(*longestPref) {
					*longestPref = currPref
				}
			}
			continue
		}

		findLongest(longestPref, currPref+string(n.val), n.next)
	}

}

func longestRepeatedPrefix(strs []string) string {

	prefTree := make(nodeIndex)

	// build out prefix trees
	for _, s := range strs {
		addString(s, prefTree)
	}

	// find the longest
	pref := ""
	findLongest(&pref, "", prefTree)

	return pref
}
