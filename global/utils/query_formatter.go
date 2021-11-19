package utils

func QueryFormatter(text string) string {
	// INPUT "name" --> OUTPUT %name%
	return "%" + text + "%"
}
