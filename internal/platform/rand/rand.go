package rand

import "math/rand"

// Between returns an int between min and max.
func Between(min, max int) int {
	return rand.Intn(max-min) + min
}
