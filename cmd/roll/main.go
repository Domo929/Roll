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
	age = flag.Bool("age", false, "Whether to roll 3d6 incl drama die")
)

const (
	dramaFmt = "%d*"
	normFmt  = "%d"
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	if len(flag.Args()) == 0 {
		log.Fatal("need to provide 'age [+/-]modifier or a list of die rolls (3d6, 2d8, etc)")
	}

	if flag.Args()[0] == "age" {
		ageGen(flag.Args())
		return
	}

	normGen(flag.Args())
}

func result(sides int) int {
	return rand.Intn(sides) + 1
}

func ageGen(args []string) error {
	modifier := "0"
	if len(args) == 2 {
		modifier = args[1]
	}
	modVal, negative, err := parseModifier(modifier)
	if err != nil {
		return err
	}

	dies := make([]int, 0, 3)
	sum := 0
	for i := 0; i < 3; i++ {
		res := result(6)

		dies = append(dies, res)
		sum += res
	}
	stuntPoints := dies[0] == dies[1] || dies[0] == dies[2] || dies[1] == dies[2]

	dieValsStr := fmt.Sprint(dies[0])
	for i := 1; i < len(dies); i++ {
		dieValsStr += fmt.Sprintf(" %d", dies[i])
	}
	fmt.Printf("Dies: %s*\n", dieValsStr)

	if negative {
		sum -= modVal
		fmt.Printf("Modifier: -%d\n", modVal)
	} else {
		sum += modVal
		fmt.Printf("Modifier: +%d\n", modVal)
	}

	fmt.Println("Total: ", sum)
	if stuntPoints {
		fmt.Printf("Generated %d stunt points\n", dies[2])
	}

	return nil
}

func parseModifier(modifier string) (int, bool, error) {
	var (
		numIndexStart = 0
		negative      = false
	)

	if len(modifier) >= 2 {
		numIndexStart = 1
		if modifier[0] == '-' {
			negative = true
		}
	}
	val, err := strconv.ParseInt(modifier[numIndexStart:], 10, 64)
	if err != nil {
		return 0, false, err
	}

	return int(val), negative, nil
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

	if len(dieGens) == 0 {
		fmt.Println("no die combos provided")
		return
	}

	fmt.Println("total: ", total)

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
