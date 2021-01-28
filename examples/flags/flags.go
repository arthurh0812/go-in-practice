package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/arthurh0812/go-in-practice/tempconv"
)

var period = flag.Duration("period", time.Second, "sleep time specification")

var temp = tempconv.CelsiusFlag("temp", 20.0, "temperature specification")

func main() {
	flag.Parse()
	fmt.Printf("sleeping for %v...", period)
	time.Sleep(period)
	fmt.Println("Wake up!")

	fmt.Printf("it is %s outside", temp)
	if temp >= 20 {
		fmt.Println("I think I'm gonna go for a walk.")
	} else {
		fmt.Println("I'd better stay inside.")
	}
}
