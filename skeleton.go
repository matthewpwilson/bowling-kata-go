package skeleton

import (
	"errors"
	"fmt"
)

// var balls []BowlingFrame
var balls []int

// type BowlingFrame struct {
// 	Score1 int
// 	Score2 int
// }

func NewGame() {
	balls = make([]int, 0)
}

func Bowl(value int) error {

	if value > 10 || value < 0 {
		return errors.New("Value must be between 0 and 10")
	}

	balls = append(balls, value)

	if value == 10 && len(balls) % 2 == 1 {
		Bowl(0)
	}
	return nil
}

func Score() int {
	var total = 0
	for i := 0; i < len(balls); i++ {
		fmt.Printf("[START] index: %d\n", i)
		if i%2 == 0 && i > 1 {
			// Detect a Strike situation
			if balls[i-2] == 10 {
				if i < len(balls)-2 && balls[i-1] == 0  {
					fmt.Printf("[Strike] Adding %d and %d to %d\n", balls[i], balls[i+1], total)
					total += balls[i] 
					total += balls[i+1]
				}
			} else {
				// Detect a Spare situation
				if (balls[i-1] + balls[i-2]) == 10 {
					fmt.Printf("[Spare] Adding %d to %d\n", balls[i], total)
					total += balls[i]
				}
			}
		}

		fmt.Printf("[END] Adding %d to %d\n", balls[i], total)
		total += balls[i]

	}
	return total
}
