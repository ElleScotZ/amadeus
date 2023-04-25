package main

import (
	"amadeus/pkg"
)

func main() {
	application := pkg.NewApplication()

	err := application.Start()
	if err != nil {
		panic(err)
	}
}
