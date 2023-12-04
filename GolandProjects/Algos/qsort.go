package main

func qsort(n int,
	less func(i, j int) bool,
	swap func(i, j int)) {
	sort1(less, swap, 0, n-1)
}

func sort1(
	less func(i, j int) bool,
	swap func(i, j int),
	left, right int) {
	if right-left+1 < 5 {
		for j := right; j > left; j-- {
			max := j
			for i := j - 1; i >= left; i-- {
				if less(max, i) {
					max = i
				}
			}
			swap(max, j)
		}
	} else {
		for left < right {
			i, j := left, left
			for j < right {
				if less(j, right) {
					swap(i, j)
					i++
				}
				j++
			}
			swap(i, right)

			if i-left < right-i {
				sort1(less, swap, left, i-1)
				left = i + 1
			} else {
				sort1(less, swap, i+1, right)
				right = i - 1
			}
		}
	}
}

/*func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	qsort(len(a),
		func(i, j int) bool { return a[i] < a[j] },
		func(i, j int) { a[i], a[j] = a[j], a[i] })

	for _, x := range a {
		fmt.Printf("%d ", x)
	}
	fmt.Println("")
}*/
