package utils

func IntConvertBool(i int) bool {
	if i == 1 {
		return true
	}
	return false
}
func BoolConvertInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
