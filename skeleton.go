package skeleton

import (
	"errors"
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

func (f *Frame) isEmpty() bool {
	return f.one == -1
}

func (f *Frame) isComplete() bool {
	return f.two != -1
}

func (f *Frame) isStrike() bool {
	return f.one == 10
}

func (f *Frame) isSpare() bool {
	return 10 == f.one+f.two
}

func (f *Frame) bowl(value int) {
	if f.isEmpty() {
		f.bowlFirst(value)
	} else if !f.isComplete() {
		f.bowlSecond(value)
	}
}

// var currentFrame *Frame
// var firstFrame *Frame
var frames []Frame

var bonusBalls Bonus

// Bonus balls
type Bonus struct {
	bonusOne int
	bonusTwo int
}

func (b *Bonus) bowl(value int) error {

	if b.secondBonusBowled() {
		return errors.New("Already bowled both bonus balls!")
	}

	if !b.firstBonusBowled() {
		if lastFrameIsStrike() || lastFrameIsSpare() {
			b.bonusOne = value
		} else {
			return errors.New("No bonus balls allowed!!")
		}
	} else {
		if lastFrameIsStrike() {
			b.bonusTwo = value
		} else {
			return errors.New("Only one bonus ball allowed!")
		}
	}

	return nil
}

func (b *Bonus) firstBonusBowled() bool {
	return b.bonusOne != -1
}

func (b *Bonus) secondBonusBowled() bool {
	return b.bonusTwo != -1
}

var frameNumber int

func NewGame() {
	frames = make([]Frame, 10)
	bonusBalls = Bonus{-1, -1}
	for i := range frames {
		frames[i].one = -1
		frames[i].two = -1
	}

	frameNumber = 0
}

func Bowl(value int) error {

	if !ballIsValid(value) {
		return errors.New("Value must be between 0 and 10")
	}

	// Bonus balls
	if lastFrameIsComplete() {
		return bonusBalls.bowl(value)
	}

	// Advance to next frame if the current one is complete
	if frames[frameNumber].isComplete() {
		frameNumber++
	}

	// Bowl the ball :)
	frames[frameNumber].bowl(value)

	return nil
}

func Score() int {
	var total = 0
	for i := 0; i < len(frames); i++ {
		if frames[i].isStrike() {
			total += getStrikeBonus(i)
		} else if frames[i].isSpare() {
			total += getSpareBonus(i)
		} else {
			// Carry on, frame was a normal (not worth 10) frame
		}
		total += frames[i].one
		total += frames[i].two
	}
	return total
}

func getSpareBonus(frameIndex int) int {
	if !isLastFrame(frameIndex) {
		return frames[frameIndex+1].one
	} else {
		// 1 bonus ball
		return bonusBalls.bonusOne
	}
}

func getStrikeBonus(frameIndex int) int {
	var total = 0
	if !isLastFrame(frameIndex) {
		total += frames[frameIndex+1].one + frames[frameIndex+1].two

		if frames[frameIndex+1].isStrike() {
			// We only added one bowl (a strike), we need to add one more
			if !isLastFrame(frameIndex + 1) {
				// We can go one frame further
				total += frames[frameIndex+2].one
			} else {
				// Strike in frame 9 and 10 - we need the first bonus ball
				total += bonusBalls.bonusOne
			}
		}
	} else {
		// Strike in last frame - we need both bonus balls
		total += bonusBalls.bonusOne + bonusBalls.bonusTwo
	}

	return total
}

func isLastFrame(i int) bool {
	return i == len(frames)-1
}

func ballIsValid(value int) bool {
	return value <= 10 && value >= 0
}

func lastFrameIsComplete() bool {
	return frames[9].isComplete()
}

func lastFrameIsSpare() bool {
	return lastFrameIsComplete() && frames[9].isSpare()
}

func lastFrameIsStrike() bool {
	return lastFrameIsComplete() && frames[9].isStrike()
}
