package internal

import (
	"crypto"
	"math"
)

func findDeltas(input []byte, blockSize int64, blockSignatures []BlockSignature) {
	weakHashToBlockIndex := make(map[string]int64)
	strongHashToBlockIndex := make(map[string]int64)

	inputLen := int64(len(input))
	a, b, s := firstRollingHash(input)

	//search s in weakhash, if its found,then calculate md5 hash and check if it matches

	for i := int64(1); (i + blockSize) < inputLen; i++ {
		currentBlock := input[i : i+blockSize]
		//search s in weakhash, if its found,then calculate md5 hash and check if it matches
		// if it matches, then store the start-end index of the block to be reused
	}
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
		start := (numberOfBlocks - 1) * blockSize
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
