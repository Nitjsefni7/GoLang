package main


type Dog interface {
	Speak()
}

type GoldenRetriever struct {
	Name  string
	Age   int
	Owner string
}

type Poodle struct {
	Name  string
	Age   int
	Owner string
}

func main() {
	golden := GoldenRetriever{Name: "Charlie", Age: 3, Owner: "Leoon"}
	poodle := Poodle{Name: "Vanessa", Age: 4, Owner: "Tom"}

}

func testSpeak(dog Dog) {
	dog.Speak()
}
func (gr GoldenRetriever) Speak() {
	fmt.Println("I am a golden retriever!")
}

func (p Poodle) Speak() {
	fmt.Println("I am a poodle!")
}