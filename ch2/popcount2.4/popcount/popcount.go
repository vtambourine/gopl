package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	c := 0
	for i := 0; i < 64; i++ {
		if x&1 > 0 {
			c++
		}
		x >>= 1
	}
	return c
}
