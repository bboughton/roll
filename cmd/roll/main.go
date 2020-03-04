package main

import (
	"fmt"
	"os"

	"github.com/bboughton/roll/pkg/dice"
)

func main() {
	args := os.Args

	for _, arg := range args[1:] {
		d, err := dice.Parse(arg)
		if err != nil {
			continue // skip errors
		}
		fmt.Println(d.Roll())
	}
}
