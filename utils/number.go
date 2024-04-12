package utils

func Page(a int, defaultValue int) int {
	if a <= 0 {
		return defaultValue
	}
	return a
}
