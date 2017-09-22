package main

import (
	"fmt"

	"github.com/drkhannah/distributed/coordinator"
)

func main() {
	ql := coordinator.NewQueueListener()
	go ql.ListenForNewSource()

	var a string
	fmt.Scan(&a)
}
