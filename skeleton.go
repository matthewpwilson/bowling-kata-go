package skeleton

import (
	"errors"
	"fmt"
)

// Frame of a bowling game
type Frame struct {
	one int
	two int
}

func (f *Frame) bowlFirst(value int) {
	f.one = value

	// If the first ball of the frame is a strike,
	// make the second a zero
	if value == 10 {
		f.two = 0
	}
}

func (f *Frame) bowlSecond(value int) {
	f.two = value
}

// var currentFrame *Frame
// var firstFrame *Frame
var frames []Frame
var bonusOne int
var bonusTwo int

var currentFrame int

func NewGame() {
	frames = make([]Frame, 10)
	for i := range frames {
		frames[i].one = -1
		frames[i].two = -1
	}
	bonusOne = -1
	bonusTwo = -1
	currentFrame = 0
}

func Bowl(value int) error {

	if !ballIsValid(value) {
		return errors.New("Value must be between 0 and 10")
	}

	// Bonus balls
	if lastFrameIsComplete() {
		if frames[currentFrame].isStrike() {
			if !firstBonusBowled() {
				bonusOne = value
			} else {
				bonusTwo = value
			}
		} else if frames[currentFrame].isSpare() {
			bonusOne = value
		} else {
			return errors.New("Game is over!")
		}
	}

	// Advance to next frame if the current one is complete
	if frameIsComplete(currentFrame) && !isLastFrame(currentFrame) {
		currentFrame++
	}

	// Bowl the ball :)
	if frameIsEmpty(currentFrame) {
		frames[currentFrame].bowlFirst(value)
	} else if !frameIsComplete(currentFrame) {
		frames[currentFrame].bowlSecond(value)
	}

	return nil
}

func Score() int {
	var total = 0
	for i := 0; i < len(frames); i++ {
		fmt.Printf("[START] index: %d\n", i)

		if frames[i].isStrike() {
			if !isLastFrame(i) {
				total += frames[i+1].one + frames[i+1].two

				if frames[i+1].isStrike() {
					total += frames[i+2].one
				}
			} else {
				// STrike in last frame
				total += bonusOne + bonusTwo
			}

		} else if frames[i].isSpare() {
			total += handleSpare(i)
		} else {
			// Carry on, frame was a normal (not worth 10) frame
		}

		fmt.Printf("[END] Adding %d and %d to %d\n", frames[i].one, frames[i].two, total)
		total += frames[i].one
		total += frames[i].two
	}
	return total
}

func handleSpare(i int) int {
	if !isLastFrame(i) {
		return frames[i+1].one
	} else {
		// 1 bonus ball
		return bonusOne
	}
}

func (f *Frame) isStrike() bool {
	return f.one == 10
}

func (f *Frame) isSpare() bool {
	return 10 == f.one+f.two
}

func isLastFrame(i int) bool {
	return i == len(frames)-1
}

func frameIsComplete(i int) bool {
	return frames[i].two != -1
}

func ballIsValid(value int) bool {
	return value <= 10 && value >= 0
}

func firstBonusBowled() bool {
	return bonusOne != -1
}

func frameIsEmpty(value int) bool {
	return frames[value].one == -1
}

func lastFrameIsComplete() bool {
	return frameIsComplete(9)
}
