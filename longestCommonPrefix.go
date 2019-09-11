package random

func findLongestPrefix(pref, s string) string {
	var n int
	if len(pref) < len(s) {
		n = len(pref)
	} else {
		n = len(s)
	}

	longPref := ""
	for i := 0; i < n; i++ {
		if pref[i] != s[i] {
			return longPref
		}
		longPref += string(pref[i])
	}
	return longPref
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	pref := strs[0]
	for _, s := range strs {
		pref = findLongestPrefix(pref, s)
		if pref == "" {
			return pref
		}
	}
	return pref
}
