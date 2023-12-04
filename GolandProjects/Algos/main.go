package main

func RuneToBin(r rune) []byte {
	if r == 0 {
		return []byte{0}
	}
	var b [32]byte
	i := 31
	for ; r > 0; r, i = r>>1, i-1 {
		if r&1 == 0 {
			b[i] = 0
		} else {
			b[i] = 1
		}
	}
	return b[i+1:]
}

func BitsToByte(a []byte) byte {
	var res byte = 0
	for _, x := range a {
		res += x
		res *= 2
	}
	res /= 2
	return res
}

func encode(utf32 []rune) []byte {
	var res []byte
	for _, x := range utf32 {
		binX := RuneToBin(x)
		k := len(binX)
		temp := make([]byte, 21, 21)
		binX = append(temp[0:21-k], binX...)
		if k <= 7 {
			res = append(res, BitsToByte(binX))
		} else if k <= 11 {
			res = append(res,
				192+BitsToByte(binX[0:15]),
				128+BitsToByte(binX[15:]))
		} else if k <= 16 {
			res = append(res,
				224+BitsToByte(binX[0:9]),
				128+BitsToByte(binX[9:15]),
				128+BitsToByte(binX[15:]))
		} else if k <= 21 {
			res = append(res,
				240+BitsToByte(binX[0:3]),
				128+BitsToByte(binX[3:9]),
				128+BitsToByte(binX[9:15]),
				128+BitsToByte(binX[15:]))
		}
	}
	return res
}

func decode(utf8 []byte) []rune {
	var res []rune
	var buf rune
	for i := 0; i < len(utf8); i++ {
		if utf8[i]/16 == 15 {
			buf = (rune)(utf8[i] % 8)
			i++
			buf = buf*64 + (rune)(utf8[i]%64)
			i++
			buf = buf*64 + (rune)(utf8[i]%64)
			i++
			buf = buf*64 + (rune)(utf8[i]%64)
			res = append(res, buf)
		} else if utf8[i]/32 == 7 {
			buf = (rune)(utf8[i] % 16)
			i++
			buf = buf*64 + (rune)(utf8[i]%64)
			i++
			buf = buf*64 + (rune)(utf8[i]%64)
			res = append(res, buf)
		} else if utf8[i]/64 == 3 {
			buf = (rune)(utf8[i] % 32)
			i++
			buf = buf*64 + (rune)(utf8[i]%64)
			res = append(res, buf)
		} else {
			buf = (rune)(utf8[i] % 128)
			res = append(res, buf)
		}

	}
	return res
}

//func main() {
/*var in string
fmt.Scan(&in)

str32 := []rune(in)

fmt.Println(RuneToBin(str32[0]))

for _, x := range decode(encode(str32)) {
	fmt.Printf("%c", x)
}
fmt.Printf("\n%s\n", encode(str32))
*/
//}
