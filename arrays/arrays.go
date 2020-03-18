package arrays

func WriteToArray(a *[2]string) {
	a[0] = "bbc"
	a[1] = "cnn"
}
func WriteToSlice(s *[]string) {
	*s = append(*s, "bbc")
	*s = append(*s, "cnn")
}
