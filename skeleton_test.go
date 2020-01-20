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

func TestBowlAllZeros(t *testing.T) {
	playGame(0)
	assert.Equal(t, 0, Score())
}

func TestBowlAllOnes(t *testing.T) {
	playGame(1)
	assert.Equal(t, 20, Score())
}

func TestBowlElevenFails(t *testing.T) {
	result := Bowl(11)
	assert.Error(t, result)
}

func TestBowlNegativeFails(t *testing.T) {
	result := Bowl(-1)
	assert.Error(t, result)
}

func TestSpareGame(t *testing.T) {
	Bowl(5)
	Bowl(5)
	Bowl(3)
	assert.Equal(t, 16, Score())
}
