package main

import (
	"math/rand"
	"fmt"
	"time"
)


func main() {
	rand.Seed(time.Now().UnixNano())

	a := rand.Intn(5)
	b := rand.Intn(5)
	fmt.Println(a, " ", b)
}