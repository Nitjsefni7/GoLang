package main

func main(){
	c1 := make(chan int)
	c2 := make(chan int)

	go listenNetwork(c1)
	go readWrite(c2)

	for {
		var data int

		select {
		case data = <-c1:
			//do stuff
		case data = <-c2:
			//do stuff
		default:
			//when none channels are ready to send information
		}
	}

	c := make(chan int)

	for num := range c{
		 //num := <-c 
		 fmt.Ptintln(num)
	}




}

func listenNetwork(c chan int){
	for {
		// read network data
		c <- data
	}
}

func readWrite(c chan int){
	for {
		//read and write to disk
		c <- data
	}
}

func doWork(c chan int) {
	for i := 0; i < 100; i++{

		c <- i
	}

	close(c)
}