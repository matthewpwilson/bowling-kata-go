package skeleton

import (
	"errors"
	"fmt"
)

var balls []int

func Bowl(value int) error {

	if value > 10 || value < 0 {
		return errors.New("Value must be between 0 and 10")
	}

	balls = append(balls, value)

	return nil
}

func Score() int {
	var total = 0
	for i := 0; i < len(balls); i++ {
		fmt.Printf("[START] index: %d\n", i)
		if i%2 == 0 {
			if (balls[i-1] + balls[i-2]) == 10 {
				fmt.Printf("[Spare] Adding %d to %d\n", balls[i], total)
				total += balls[i]
			}
		}

		fmt.Printf("[END] Adding %d to %d\n", balls[i], total)
		total += balls[i]

	}
	return total
}
