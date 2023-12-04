package main

func calculate(a []byte) int {
	var stack []int
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '+' {
			k := len(stack)
			stack = append(stack[0:k-2], stack[k-1]+stack[k-2])
		} else if a[i] == '*' {
			k := len(stack)
			stack = append(stack[0:k-2], stack[k-1]*stack[k-2])
		} else if a[i] == '-' {
			k := len(stack)
			stack = append(stack[0:k-2], stack[k-1]-stack[k-2])
		} else if a[i] <= '9' && a[i] >= '0' {
			stack = append(stack, (int)(a[i]-48))
		}
	}
	return stack[0]
}

/*func main() {
	reader := bufio.NewReader(os.Stdin)

	in, _ := reader.ReadString('\n')

	a := []byte(in)
	res := calculate(a)
	fmt.Println(res)
}*/
