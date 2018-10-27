package main

import (
	"fmt"
	"os"

	"github.com/martinusso/inflect"
)

func main() {
	fmt.Println(inflect.Pluralize(os.Args[1]))
}
