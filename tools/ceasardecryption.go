package tools

func CeasarDecryptionWithShift(shift int, encryptedtext string) string {
	var letterascii int
	var rawtext string

	for i := 0; i < len(encryptedtext); i++ {
		letterascii = int(encryptedtext[i])

		letterascii -= shift

		if letterascii < 97 {
			letterascii = letterascii + 26 // 123 - (97 - letterascii), basically finding the ascii undercount and subtracting it from the end
		}

		rawtext += string(rune(letterascii)) // beter way of doing this??
	}

	return rawtext
}

func CeasarDecryptionWithoutShift(encryptedtext string) []string {
	var rawtext string
	var finaldecryptions []string

	for i := 0; i < 27; i++ {
		rawtext = CeasarDecryptionWithShift(i, encryptedtext)
		finaldecryptions = append(finaldecryptions, rawtext)
	}

	return finaldecryptions
}
