package main

import "log"

func main() {

	a := 10
	b := 5

	if a-b > 3 {
		log.Println("True")
	}

	if a*5 == 50 {
		panic(1)
	}
}
