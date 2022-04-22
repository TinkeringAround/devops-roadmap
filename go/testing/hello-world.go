package main

func Hello(you, lang string) string {
	var name = you
	var prefix string

	switch lang {
	case "Spanish":
		prefix = "Hola, "
	case "French":
		prefix = "Bonjour, "
	default:
		prefix = "Hello, "
	}

	if you == "" {
		name = "World"
	}

	return prefix + name
}
