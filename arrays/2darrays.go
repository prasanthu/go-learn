package arrays

func New2DSlice(x, y int) [][]int {
	var a = make([][]int, y)
	var backing = make([]int, x*y)
	for i := range a {
		a[i], backing = backing[:x], backing[x:]
	}
	return a
}

func IdentityFill(a [][]int) {
	for i := range a {
		for j := range a[i] {
			a[i][j] = 0
		}
		a[i][i] = 1
	}
}
