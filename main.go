package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	sum := add(1, 2)

	var x int
	x = 3
	var b uint

	a := 1
	f := float32(a)

	const c = 1

	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	// w Go nie ma petli while, więc przerabia się pętle for, żeby działała podobnie
	for sum2 < 10 {
		sum++
	}
	//nieskonczona pętla
	for {
		sum++

		if sum == 5 {
			continue
		} else if sum > 20 {
			break
		} else {
			// whatever
		}
		// whatever
	}

	//przy if mozemy od razu inicjalizowac zmienne
	if sum := nazwa_funkcji(); sum < 10 {
		//do smth
	}

	//switch - przy poszczególnych case nie potrzeba break, poniewaz w Go po jakimkolwiek casie cały switch zostanie zakończony
	sum := nazwa_funkcji()

	switch sum {
	case 50:
		// do smth
	case 100:
		//cos tam
	default:
		//do smth else
	}

	//Go specific -> defer opóźnia wykonanie kodu dopóki funkcja w ktorej się nie znajduje nie zakończy działania
	//defet odkładane są na stack, więc jeśli jest ich więcej zostaną wykonane po zakończeniu od tyłu (od ostatniego defer)
	defer fmt.Println("Hello World")

	sum := 1
	fmt.Println(sum)
}

func add(x int, y int) int {
	return x + y
}