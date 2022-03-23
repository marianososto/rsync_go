package internal

import (
	"crypto"
	"math"
)

const mod = 1 << 16

func rollingHash(block []byte) (r1 uint32, r2 uint32, r uint32) {
	var a, b uint32
	l := uint32(len(block))
	for i := uint32(0); i < l; i++ {
		a += uint32(block[i])
		b += (l - i + 1) * uint32(block[i])

	}
	r1 = a % mod
	r2 = b % mod
	r = r1 + mod*r2
	return r1, r2, r
}

type BlockSignature struct {
	blockIndex int64
	blockSize  int64
	weakHash   uint32
	strongHash []byte
}

func processDestinationFile(input []byte) []BlockSignature {

	inputLen := int64(len(input))
	blockSize := int64(math.Floor(math.Sqrt(float64(inputLen))))
	numberOfBlocks := int64(math.Ceil(float64(int64(inputLen) / blockSize)))

	blocksSignatures := make([]BlockSignature, 0, numberOfBlocks)

	for i := int64(0); i < (numberOfBlocks - 1); i++ {
		start := i * blockSize
		end := start + blockSize
		block := input[start:end]
		_, _, weakHash := rollingHash(block)
		strongHash := crypto.MD5.New().Sum(block)
		blocksgn := BlockSignature{
			blockSize:  blockSize,
			blockIndex: i,
			strongHash: strongHash,
			weakHash:   weakHash,
		}
		blocksSignatures = append(blocksSignatures, blocksgn)
	}

	if inputLen-((numberOfBlocks-1)*blockSize) > 0 {
		start := numberOfBlocks * blockSize
		block := input[start:]
		_, _, weakHash := rollingHash(block)
		strongHash := crypto.MD5.New().Sum(block)
		blocksgn := BlockSignature{
			blockSize:  inputLen - ((numberOfBlocks - 1) * blockSize), // last block is not fixed, it's the remainder till the end of the array
			blockIndex: numberOfBlocks - 1,
			strongHash: strongHash,
			weakHash:   weakHash,
		}
		blocksSignatures = append(blocksSignatures, blocksgn)
	}
	return blocksSignatures
}
