package sort

func Bubble(x []int) {
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
}


func Dac(x []int) []int {

	if len(x) < 100 {
		Bubble(x)
		return x
	}

	// get min and max
	min := x[0]
	max := x[0]
	for _, v := range x {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	// divide
	center := max - ((max - min) / 2)

	below := make([]int, 0)
	above := make([]int, 0)
	for _, v := range x {
		if v < center {
			below = append(below, v)
		} else {
			above = append(above, v)
		}
	}

	below = Dac(below)
	above = Dac(above)
	below = append(below, above...)

	// recursion
	return below
}