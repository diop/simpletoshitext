package main

import (
	"fmt"

	"github.com/diop/simpletoshitext/cmd/simpletoshitext/router"
)

func main() {
	fmt.Println("Live from the server...")

	e := router.New()

	e.Start(":8000")
}
