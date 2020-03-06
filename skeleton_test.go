package skeleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func playGame(pins int) error {
	for i := 0; i < 20; i++ {
		error := Bowl(pins)
		if error != nil {
			return error
		}
	}
	return nil
}

func bowlMany(value int, of int) {
	for i := 0; i < of; i++ {
		Bowl(value)
	}
}

func TestBowlAllZeros(t *testing.T) {
	NewGame()
	playGame(0)
	assert.Equal(t, 0, Score())
}

func TestBowlAllOnes(t *testing.T) {
	NewGame()
	playGame(1)
	assert.Equal(t, 20, Score())
}

func TestBowlElevenFails(t *testing.T) {
	NewGame()
	result := Bowl(11)
	assert.Error(t, result)
}

func TestBowlNegativeFails(t *testing.T) {
	NewGame()
	result := Bowl(-1)
	assert.Error(t, result)
}

func TestSpareGame(t *testing.T) {
	NewGame()
	Bowl(5)
	Bowl(5)
	Bowl(3)
	bowlMany(0, 17)

	assert.Equal(t, 16, Score())
}

func TestSpareInLastFrame(t *testing.T) {
	NewGame()
	bowlMany(0, 18)
	Bowl(5)
	Bowl(5)

	Bowl(1) // bonus ball

	assert.Equal(t, 11, Score())
}

func TestBonusFailsIfNoSpare(t *testing.T) {
	NewGame()
	bowlMany(0, 18)
	Bowl(5)
	Bowl(4)

	result := Bowl(1) // bonus ball
	assert.Error(t, result)
}

func TestStrike(t *testing.T) {
	NewGame()
	Bowl(10)
	Bowl(5)
	Bowl(3)
	bowlMany(0, 16)

	assert.Equal(t, 26, Score())
}

func TestStrikeInLastFrame(t *testing.T) {
	NewGame()
	bowlMany(0, 18)
	Bowl(10)

	Bowl(5) // bonus ball 1
	Bowl(1) // bonus ball 2

	assert.Equal(t, 16, Score())
}

func TestTurkey(t *testing.T) {
	NewGame()
	Bowl(10) // 30
	Bowl(10) // 25
	Bowl(10) // 19
	Bowl(5)  // 5
	Bowl(4)  // 4
	bowlMany(0, 12)
	assert.Equal(t, 83, Score())
}
