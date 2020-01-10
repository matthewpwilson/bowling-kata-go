package skeleton

import "errors"

var total int

func Bowl(value int) error {
	if value > 10 || value < 0 {
		return errors.New("Value must be between 0 and 10")
	}
	total += value
	return nil
}

func Score() int {
	return total
}
