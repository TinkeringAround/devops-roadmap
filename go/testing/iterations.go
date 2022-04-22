package main

func Repeat(character string, times int) string {
	var repeatedCharacter = character

	for i := 0; i < (times - 1); i++ {
		repeatedCharacter += character
	}

	return repeatedCharacter
}
