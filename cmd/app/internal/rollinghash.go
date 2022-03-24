package internal

const mod = 1 << 16

func rollingHash(block []byte, blockSize uint32, start, prevA, prevB uint32) (a uint32, b uint32, s uint32) {
	end := start + blockSize
	a = (prevA - uint32(block[start-1]) + uint32(block[start+blockSize])) % mod
	b = (prevB - (end-start+1)*uint32(block[start]) + a) % mod
	return a, b, a + mod*b
}

func firstRollingHash(block []byte) (a uint32, b uint32, s uint32) {
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
