package main

import "fmt"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

func HelloWord() string {
	return englishPrefix + "World"
}

func HelloWordWithName(name string) string {
	if name == "" {
		name = "World"
	}
	return englishPrefix + name
}

func HelloWorldWithOtherLanguage(language, name string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(HelloWord())
	fmt.Println(HelloWordWithName("ikan"))
	fmt.Println(HelloWordWithName(""))
	fmt.Println(HelloWorldWithOtherLanguage("Spanish", ""))
	fmt.Println(HelloWorldWithOtherLanguage("Spanish", "Leo"))
	fmt.Println(HelloWorldWithOtherLanguage("French", "Leo"))
	fmt.Println(HelloWorldWithOtherLanguage("", "Leo"))
	fmt.Println(HelloWorldWithOtherLanguage("French", ""))
	fmt.Println(HelloWorldWithOtherLanguage("", ""))
}
