package main

import (
	"fmt"

	"github.com/acidleroy/elevation/server"
)

func main() {
	fmt.Println("Starting the elevation server")
	server.Start()
}
