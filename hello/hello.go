package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// for i := 0; i < 10; i++ {
	// 	message, err := greetings.Hello("Grz")
	

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println(message)
	// }

	names := []string{"Grz", "Monia", "Paszczak", "Watson", "Pajda"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for i, v := range messages {
		fmt.Println(i + ": " + v)
	}
}