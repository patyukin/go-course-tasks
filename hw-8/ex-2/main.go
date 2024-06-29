package main

func MinEl(a []int) int {
	if len(a) == 0 {
		return 0
	}

	if len(a) == 1 {
		return a[0]
	}

	t := MinEl(a[:len(a)-1])
	if t <= a[len(a)-1] {
		return t
	}

	return a[len(a)-1]
}

func MinElRecLoop(a []int) int {
	if len(a) == 0 {
		return 0
	}

	minEl := a[0]
	for _, v := range a {
		if v < minEl {
			minEl = v
		}
	}

	return minEl
}
