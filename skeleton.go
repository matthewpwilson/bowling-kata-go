package skeleton

import (
	"errors"
	"fmt"
)

type Frame struct {
	one int
	two int
}

var balls []int

func NewGame() {
	balls = make([]int, 0)
}

func Bowl(value int) error {

	if value > 10 || value < 0 {
		return errors.New("Value must be between 0 and 10")
	}

	balls = append(balls, value)

	if value == 10 && len(balls)%2 == 1 {
		Bowl(0)
	}
	return nil
}

func Score() int {
	var total = 0
	for i := 0; i < len(balls); i++ {
		fmt.Printf("[START] index: %d\n", i)
		if isFrameBoundary(i) {
			if isTurkey(i) { // We think we might want to revert this start of a change,
				// and do the frames refactor next.

			} else if isStrike(i) {
				if i < len(balls)-1 {
					fmt.Printf("[Strike] Adding %d and %d to %d\n", balls[i], balls[i+1], total)
					total += balls[i]
					total += balls[i+1]
				} //this logic might not work for consecutive strikes
			} else {
				// Detect a Spare situation
				if isSpare(i) {
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

func isFrameBoundary(i int) bool {
	return i%2 == 0 && i > 1
}

func isStrike(i int) bool {
	return balls[i-2] == 10 //&& balls[i-1] == 0
}

func isSpare(i int) bool {
	return (balls[i-1] + balls[i-2]) == 10
}

func isTurkey(i int) bool {
	return balls[i-6] == 10 && balls[i-4] == 10 && isStrike(i)
}
