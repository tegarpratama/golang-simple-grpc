package main

import (
	"project-two/cmd"
)

func main() {
	go cmd.ServeGRPC()
	cmd.ServeHTTP()
}
