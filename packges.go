package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("my favorite number is ", rand.Intn(10))
	fmt.Println("now you %g nums", rand.Intn(10))
	fmt.Println("g %s", getName("tarou"))
}

func getName(name string) string {
	return name
}
