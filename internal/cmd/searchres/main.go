package main

import (
	"fmt"

	"tonek.org/tenn/internal/activecomm"
)

func main() {
	cl := activecomm.NewClient()
	res, err := cl.FindResources("wood")
	check(err)
	fmt.Printf("%+v\n", res)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
