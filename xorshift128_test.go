package xorshift128

import (
	"testing"
)

var xorshift128Tests = []struct {
	seed0, seed1, expected uint64
}{
	{1, 2, 8388677},
	{3, 4, 25166027},
	{79, 480, 662705231},
	{47, 9289, 394286063},
}

//TestXorshift128 tests the xorshift128+ algorithm with certain seeds
func TestXorshift128(t *testing.T) {
	for _, tt := range xorshift128Tests {
		actual := Num(tt.seed0, tt.seed1)
		if actual != tt.expected {
			t.Errorf("Calc(%d, %d): expected %d, actual %d", tt.seed0, tt.seed1, tt.expected, actual)
		}
	}
}
