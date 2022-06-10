package main

import "fmt"

func main() {

	result := divider(1, 0)
	fmt.Println("결과는:", result)

}

func divider(a int, b int) int {
	response := a / b
	return response
}
