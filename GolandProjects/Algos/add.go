package main

func add(a, b []int32, p int) []int32 {
	var maxLen int
	if len(b) > len(a) {
		maxLen = len(b)
	} else {
		maxLen = len(a)
	}

	res := make([]int32, maxLen+1)
	var dx int32 = 0
	for i := 0; i < maxLen; i++ {
		sum := dx
		if i < len(a) {
			sum += a[i]
		}
		if i < len(b) {
			sum += b[i]
		}

		res[i] = sum % (int32)(p)
		dx = sum / (int32)(p)
	}
	if dx > 0 {
		res[maxLen] = dx
	} else {
		res = res[:maxLen]
	}

	return res
}

/*func main() {
	fmt.Println(28 * 0.6)
}*/
