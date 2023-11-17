package main

func main() {
	var c1 = make(chan int)
	var c2 = make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			if i%2 == 0 {
				c1 <- i
			} else {
				c2 <- i
			}
		}
	}()

	for {
		select {
		case v := <-c1:
			println("Received from c1:", v)
		case v := <-c2:
			println("Received from c2:", v)
		}
	}

}
