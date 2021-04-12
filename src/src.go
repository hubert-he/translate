package lang

func Translate(msg string) (hello string) {
	switch msg {
	case "jp":
		hello = "こんにちは、Go言語"
	case "cn":
		hello = "你好，Go语言"
	case "fr":
		hello = "Bonjour, Go language"
	case "de":
		return "Hallo, geh Sprache"
	case "es":
		hello = "Hola, ir idioma"
	default:
		hello = "Hello, Go Language"
	}
	return
}
