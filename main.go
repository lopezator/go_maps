package main

import "fmt"

func main() {
	//var colors map[string]string
	//colors := make(map[string]string)

	colors := map[string]string{
		"red": "rojo",
		"green": "verde",
		"blue": "azul",
	}

	delete(colors, "green")

	fmt.Println(colors)
}
