package rolls

import (
	"fmt"
	"strconv"
)

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
