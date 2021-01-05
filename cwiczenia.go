//funkcja wypisujaca slowa ze stringa i ich ilosc

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for _, value := range words {
		if elem, ok := m[value]; ok {
			m[value] = elem + 1
		} else {
			m[value] = 1
		}
	}
	return m
}

//closure

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

//fibonacci closure
func fibonacci() func() int {
	f2 := 0
	f := 1
	return func() int {
		f, f2 = f2, f+f2
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}