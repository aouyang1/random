package random

func defangIPaddr(address string) string {
	var defanged string
	for _, v := range address {
		if v == '.' {
			defanged += "[.]"
		} else {
			defanged += string(v)
		}
	}
	return defanged
}
