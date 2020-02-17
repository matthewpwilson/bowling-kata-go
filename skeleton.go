package skeleton

import (
	"errors"
	"fmt"
)

// Frame of a bowling game
type Frame struct {
	one   int
	two   int
	bonus int
}

// var currentFrame *Frame
// var firstFrame *Frame
var frames []Frame

var currentFrame int

func NewGame() {
	frames = make([]Frame, 10)
	for i := range frames {
		fmt.Printf("Setting up frame %d\n", i)
		frames[i].one = -1
		frames[i].two = -1
		frames[i].bonus = 0
	}
	currentFrame = 0
}

func Bowl(value int) error {

	if value > 10 || value < 0 {
		return errors.New("Value must be between 0 and 10")
	}

	// Advance to next frame if the current one is complete
	if frames[currentFrame].two != -1 {
		currentFrame++
	}

	if frames[currentFrame].one == -1 {
		frames[currentFrame].one = value
		// Strike
		if value == 10 {
			frames[currentFrame].two = 0
		}
	} else if frames[currentFrame].two == -1 {
		frames[currentFrame].two = value
	}

	return nil
}

func Score() int {
	var total = 0
	for i := 0; i < len(frames); i++ {
		fmt.Printf("[START] index: %d\n", i)

		if frameIsStrike(i) {
			// Got a strike
			if isLastFrame(i) {
				continue
			}

			frames[i].bonus = frames[i+1].one + frames[i+1].two
		} else if frameIsSpare(i) {
			// Got a spare
			if isLastFrame(i) {
				continue
			}

			frames[i].bonus = frames[i+1].one
		} else {
			// Carry on
		}

		fmt.Printf("[END] Adding %d and %d to %d\n", frames[i].one, frames[i].two, total)
		total += frames[i].one
		total += frames[i].two
		total += frames[i].bonus
	}
	return total
}

func frameIsStrike(i int) bool {
	return frames[i].one == 10
}

func frameIsSpare(i int) bool {
	return 10 == frames[i].one+frames[i].two
}

func isLastFrame(i int) bool {
	return i == len(frames)
}
