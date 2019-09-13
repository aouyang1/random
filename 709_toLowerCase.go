package random

func toLowerCase(str string) string {
	var out string
	for _, s := range str {
		if 'A' <= s && s <= 'Z' {
			s += 'a' - 'A'
		}
		out += string(byte(s))
	}
	return out
}
