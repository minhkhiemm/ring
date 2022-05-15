package main

import (
	"fmt"

	"github.com/minhkhiemm/ring/broadcast"
)

func main() {
	fmt.Println("Hi")

	broadcast.Broadcast(10, 3, 4, 2, 1)
}
