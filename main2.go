package main

import "fmt"

//struct
type Dog struct {
	Name 	string
	Age 	int
	Owner 	string
}


func main() {
	//pointers 
	num := 1

	numPtr := &num
	*numPtr = 3

	//struct
	var dog Dog

	dog.Name = "Charlie"
	dog.Age = 3
	dog.Owner = "Grz"

	fmt.Println(dog)

	dog := Dog{Name: "Charlie", Age: 3, Owner: "Grz"}

	//arrays
	//zmiany w slice zmieniaja sie rowniez w oryginalnej tablicy
	var array [10]int

	array[0] = 1
	array[1] = 2

	array := [10]int{1, 2, 3, 4, 5, 6}
	slice := array[0:3]

	fmt.Println(array[0])
	//pracowanie na arrays ale przy uzyciu slice poniewaz mozna np. zmieniac ich rozmiar
	slice := []int{}
	slice = append(slice, 1, 2, 3, 4, 5)

	//petla do slice, i to index i num to wartosc pod tym indeksem
	for i, num := range.slice {
		fmt.Println(i, num)
	}

	//maps
	var m map[string]int
	m = make(map[string]int)
	// krotsza wersja
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2

	delete(m, "two")

	val, ok := m["three"]

	fmt.Println(m["one"])
}