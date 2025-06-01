package hello_world

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "hello, "
	spanishHelloPrefix = "hola, "
	frenchHelloPrefix  = "bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	prefix := ""

	switch language {
	case "":
		prefix = englishHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}
