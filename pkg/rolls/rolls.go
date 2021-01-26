package rolls

import (
	"flag"
	"log"
	"math/rand"
)

func Roll(args []string) {
	if args[0] == "age" {
		if err := ageGen(flag.Args()); err != nil {
			log.Println(err)
		}
		return
	}

	normGen(args)
}

func result(sides int) int {
	return rand.Intn(sides) + 1
}
