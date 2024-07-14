package libs

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestProof(t *testing.T) {
	leaves := []common.Hash{
		common.BytesToHash([]byte("dataHash1")),
		common.BytesToHash([]byte("dataHash2")),
		common.BytesToHash([]byte("dataHash3")),
		common.BytesToHash([]byte("dataHash4")),
	}
	root := CalculateMerkleRoot(leaves)
	proof := GenerateMerkleProof(leaves, 0)
	fmt.Println(root.String())
	for i := 0; i < len(proof); i++ {
		fmt.Println(proof[i].String())
	}
	fmt.Println(hashPair(proof[0], proof[1]))
}
