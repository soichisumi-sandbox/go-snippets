package main

import "fmt"

func channelBasic() {
	// closeFirst
	c := make(chan int, 4)
	for _, v := range []int{ 1, 4, 25, 52 } {
		c <- v
	}
	close(c)
	for v := range c {
		fmt.Printf("%d received\n", v)
	}
}

func main(){
	channelBasic()
}