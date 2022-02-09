package tools

func CeasarEncryption(shift int, rawtext string) string {
	var letterascii int
	var encryptedtext string

	for i := 0; i < len(rawtext); i++ {
		letterascii = int(rawtext[i])

		letterascii += shift

		if letterascii > 122 {
			letterascii = letterascii - 25 // letterascii - 122 + 97, basically finding the ascii overcount and adding it back from the beginning
		}

		encryptedtext += string(rune(letterascii)) // beter way of doing this??
	}

	return encryptedtext
}
