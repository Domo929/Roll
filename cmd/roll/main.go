package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	race = flag.Bool("race", false, "Whether to roll 3d6 incl drama die")
)

const (
	dramaFmt = "%d *\n"
	normFmt  = "%d\n"
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	if *race {
		raceGen()
		return
	}

	normGen(flag.Args())
}

func result(sides int) int {
	return rand.Intn(sides) + 1
}

func raceGen() {
	total := 0
	for i := 0; i < 3; i++ {
		die := result(6)
		total += die

		dieFmt := normFmt
		if i == 0 {
			dieFmt = dramaFmt
		}
		fmt.Printf(dieFmt, die)
	}
	fmt.Printf("Total: %d\n", total)
}

func normGen(dieGens []string) {
	total := 0
	for _, dieGen := range dieGens {
		num, sides, err := parse(dieGen)
		if err != nil {
			log.Println(err)
			continue
		}
		resMsg := fmt.Sprintf("%s: ", dieGen)
		for i := 0; i < num; i++ {
			res := result(sides)
			total += res
			resMsg = fmt.Sprintf("%s %d", resMsg, res)
		}
		fmt.Println(resMsg)
	}
	if len(dieGens) > 0 {
		fmt.Println("total: ", total)
	} else {
		fmt.Println("no die combos provided")
	}

}

func parse(dieGen string) (int, int, error) {
	parts := strings.Split(dieGen, "d")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("passed illegal die command: %s", dieGen)
	}

	num, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, 0, err
	}

	sides, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, 0, err
	}

	return int(num), int(sides), nil
}
