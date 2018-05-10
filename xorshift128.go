package xorshift128

//Num generates a random number based off the current time as a seed
func Num(seed0, seed1 uint64) uint64 {
	s1 := seed0
	s0 := seed1
	seed0 = s0
	s1 ^= (s1 << 23)
	seed1 = (s1 ^ s0 ^ (s1 >> 17) ^ (s0 >> 26))
	return seed0 + seed1
}
