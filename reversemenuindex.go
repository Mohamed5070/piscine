package piscine

func ReverseMenuIndex(menu []string) []string {
	slice := []string{}
	for i := len(menu) - 1; i >= 0; i-- {
		slice = append(slice, menu[i])
	}
	return slice
}