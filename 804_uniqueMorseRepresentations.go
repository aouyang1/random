package random

var morseCode = map[rune]string{
	'a': ".-",
	'b': "-...",
	'c': "-.-.",
	'd': "-..",
	'e': ".",
	'f': "..-.",
	'g': "--.",
	'h': "....",
	'i': "..",
	'j': ".---",
	'k': "-.-",
	'l': ".-..",
	'm': "--",
	'n': "-.",
	'o': "---",
	'p': ".--.",
	'q': "--.-",
	'r': ".-.",
	's': "...",
	't': "-",
	'u': "..-",
	'v': "...-",
	'w': ".--",
	'x': "-..-",
	'y': "-.--",
	'z': "--..",
}

func uniqueMorseRepresentations(words []string) int {
	umr := make(map[string]struct{})
	var code string
	for _, w := range words {
		code = ""
		for _, letter := range w {
			code += morseCode[letter]
		}
		if _, exists := umr[code]; !exists {
			umr[code] = struct{}{}
		}
	}

	return len(umr)
}
