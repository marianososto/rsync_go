package internal

const mod = 1 << 16

func rollingHash(block []byte) (a uint32, b uint32, s uint32) {
	var a1, b1 uint32
	l := uint32(len(block))
	for i := uint32(0); i < l; i++ {
		a1 += uint32(block[i])
		b1 += (l - i + 1) * uint32(block[i])

	}
	a = a1 % mod
	b = b1 % mod
	s = a + mod*b
	return a, b, s
}
func rollingHashSuccessive(block []byte, start, end int64) (a uint32, b uint32, s uint32) {

	if start == 0 {
		return rollingHash(block)
	}

	prevA, prevB, _ := rollingHash(block[start-1 : end-1]) // we have to store previous values to leverage the previuous calculated value
	a = (prevA - uint32(block[start-1]) + uint32(block[end])) % mod
	b = (prevB - uint32(end-start+1)*uint32(block[start]) + a) % mod
	return a, b, a + mod*b
}
