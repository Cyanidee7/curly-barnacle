package main

import (
	"fmt"
)

var build = []string{"BoD", "DHS", "WoN", "Scarlet", "Haas"}

func main() {
	var build = append(build, "Winter", "Immo", "War-Axe")

	fmt.Println(build, len(build))
}
