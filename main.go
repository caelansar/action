package main

import (
	"fmt"

	"github.com/caelansar/private"
)

func main() {
	// private package is in a private repository
	p := private.Private[int]{
		Id: 1,
	}
	v := p.Get()
	fmt.Println(v)
}
