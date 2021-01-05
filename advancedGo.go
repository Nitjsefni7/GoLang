package main

type Dog struct {
	Name 	string
	Age 	int
	Owner 	string
}

func main() {
	dog := Dog{Name: "Alex", Age: 3, Owner: "Me"}
	//uzywamy adresu obiektu, a nie samego obiektu, gdyz uzycie obiektu w metodzie powoduje skopiowanie calego obiektu i dokonanie
	//zmian wyłącznie na tej kopii. Oryginal zostaje taki sam
		
	// dogPtr := &dog
	// dogPtr.ChangeName("Bob")
	(&dog).ChangeName("Bob")
}

func (d *Dog) ChangeName(newName strign){
	(*d).Name = newName
}