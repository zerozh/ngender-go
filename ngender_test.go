package ngender

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGuess(t *testing.T) {
	tolerance := .000001
	gender, probability := Guess("赵本山")
	assert.Equal(t, "male", gender)
	assert.True(t, math.Abs(probability-0.983622) < tolerance)

	gender, probability = Guess("宋丹丹")
	assert.Equal(t, "female", gender)
	assert.True(t, math.Abs(probability-0.975948) < tolerance)

	gender, probability = Guess("诸葛亮")
	assert.Equal(t, "male", gender)

	gender, probability = Guess("欧阳锋")
	assert.Equal(t, "male", gender)

	gender, probability = Guess("James")
	assert.Equal(t, "unknown", gender)

	gender, probability = Guess("Tinky Winky")
	assert.Equal(t, "unknown", gender)
}

func BenchmarkGuess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Guess("王伟")
	}
}
