package main

import "fmt"

func main()  {
	var first int
	var second int
	_, err := fmt.Scanf("%d %d", &first, &second)

	if err != nil {
		fmt.Println("ERROR!!")
	} else {
		sum := first + second
		fmt.Printf("%d + %d = %d \n", first, second, sum)
	}
}