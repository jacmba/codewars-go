package tortoiseracing

// Race return time array of reaching calculation
func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}

	var h, m, s int
	delta := v2 - v1

	time := int((float64(g) / float64(delta)) * 3600)
	h = time / 3600
	m = (time - (h * 3600)) / 60
	s = time - (m * 60) - (h * 3600)

	return [3]int{h, m, s}
}
