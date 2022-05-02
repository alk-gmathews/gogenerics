package arithmetic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumInts(t *testing.T) {
	m := map[string]int64{
		"first":  10,
		"second": 20,
	}

	t.Run(
		"summing ints should return 30",
		func(t *testing.T) {
			assert.Equal(t, SumInts(m), int64(30))
		})
}

func TestSumFloats(t *testing.T) {
	m := map[string]float64{
		"first":  10.60,
		"second": 20.10,
	}

	t.Run(
		"summing floats should return 10.60+20.10",
		func(t *testing.T) {
			assert.Equal(t, SumFloats(m), float64(10.60)+float64(20.10))
		})
}

func TestSumNumbers(t *testing.T) {
	t.Run(
		"adding ints should return 30",
		func(t *testing.T) {
			m := map[string]int64{
				"first":  10,
				"second": 20,
			}
			assert.Equal(t, SumNumbers(m), int64(30))
		})

	t.Run(
		"adding floats should return 10.60+20.10",
		func(t *testing.T) {
			m := map[string]float64{
				"first":  10.60,
				"second": 20.10,
			}
			assert.Equal(t, SumNumbers(m), float64(10.60)+float64(20.10))
		})
}
