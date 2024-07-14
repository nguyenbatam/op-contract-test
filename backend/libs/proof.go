package libs

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
)

func CalculateMerkleRoot(leaves []common.Hash) common.Hash {
	if len(leaves) == 0 {
		return [32]byte{}
	}

	if len(leaves) == 1 {
		return leaves[0]
	}

	var hashes []common.Hash
	for i := 0; i < len(leaves); i += 2 {
		var pair common.Hash
		if i+1 < len(leaves) {
			pair = hashPair(leaves[i], leaves[i+1])
		} else {
			pair = hashPair(leaves[i], leaves[i])
		}
		hashes = append(hashes, pair)
	}

	return CalculateMerkleRoot(hashes)
}

func GenerateMerkleProof(leaves []common.Hash, index int) []common.Hash {
	var proof []common.Hash
	for i := 0; i < len(leaves); i += 2 {
		if i+1 < len(leaves) {
			if i == index || i+1 == index {
				proof = append(proof, hashPair(leaves[i], leaves[i+1]))
			} else {
				proof = append(proof, hashPair(leaves[i], leaves[i+1]))
			}
		} else {
			proof = append(proof, hashPair(leaves[i], leaves[i]))
		}
		index = index / 2
	}
	return proof
}

func hashPair(a, b common.Hash) common.Hash {
	var pair []byte
	pair = append(pair, a[:]...)
	pair = append(pair, b[:]...)
	return sha256.Sum256(pair[:])
}
