package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/Domo929/roll/pkg/rolls"
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	if len(flag.Args()) == 0 {
		log.Fatal("need to provide 'age [+/-]modifier or a list of die rolls (3d6, 2d8, etc)")
	}

	rolls.Roll(flag.Args())
}
