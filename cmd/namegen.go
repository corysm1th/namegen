package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	namegen "github.com/corysm1th/namegen/pkg"
)

func main() {
	var n int
	var err error
	if len(os.Args) < 2 {
		n = 1
	} else {
		count := os.Args[1]
		n, err = strconv.Atoi(count)
		if err != nil {
			log.Fatalf("%s is not a valid integer", count)
		}
	}

	textCase := namegen.Case(namegen.Snake)

	for i := 0; i < n; i++ {
		fmt.Println(namegen.Things(textCase))
	}
}
