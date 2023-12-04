package main

import "fmt"

func calcOp(in string) int {
	str := []byte(in)
	used := make(map[string]int)
	calcID := 0
	var indexes []int

	for i, x := range str {
		if x == '(' {
			indexes = append(indexes, i)
		} else if x == ')' {
			expr := string(str[indexes[len(indexes)-1] : i+1])
			indexes = indexes[:len(indexes)-1]
			_, ok := used[expr]
			if !ok {
				used[expr] = calcID
				calcID++
			}
		}
	}
	return len(used)
}

func main() {
	var in string
	fmt.Scan(&in)

	fmt.Println(calcOp(in))
}
