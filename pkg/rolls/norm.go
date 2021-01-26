package rolls

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func normGen(dieGens []string) {
	total := 0
	for _, dieGen := range dieGens {
		num, sides, err := parseNormDice(dieGen)
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

func parseNormDice(dieGen string) (int, int, error) {
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
