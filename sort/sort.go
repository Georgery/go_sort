package sort

func Bubble(x []int) []int {
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < len(x); i++ {
			// fmt.Println(i)
			if x[i-1] > x[i] {
				x[i-1], x[i] = x[i], x[i-1]
				swapped = true
			}
		}
	}
	return x
}