package distributor

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestBalanceTree(t *testing.T) {
	balances := []Balance{
		{Account: common.HexToAddress("0x17ec8597ff92C3F44523bDc65BF0f1bE632917ff"), Amount: big.NewInt(200)},
		{Account: common.HexToAddress("0x63FC2aD3d021a4D7e64323529a55a9442C444dA0"), Amount: big.NewInt(300)},
		{Account: common.HexToAddress("0xD1D84F0e28D6fedF03c73151f98dF95139700aa7"), Amount: big.NewInt(250)},
	}
	tree, err := NewBalanceTree(balances)
	assert.Nil(t, err)

	root := tree.GetRoot()
	assert.Equal(t, root.Hex(), "0x2ec9c2fc2a55df417ba88ecd833f165fa3c5941772ebaf8c5f4debe33f4d1b12")

	proofs := [][]string{
		{"2a411ed78501edb696adca9e41e78d8256b61cfac45612fa0434d7cf87d916c6"},
		{"bfeb956a3b705056020a3b64c540bff700c0f6c96c55c0a5fcab57124cb36f7b", "d31de46890d4a77baeebddbd77bf73b5c626397b73ee8c69b51efe4c9a5a72fa"},
		{"ceaacce7533111e902cc548e961d77b23a4d8cd073c6b68ccf55c62bd47fc36b", "d31de46890d4a77baeebddbd77bf73b5c626397b73ee8c69b51efe4c9a5a72fa"},
	}
	for idx, balance := range balances {
		p := make(Elements, 0)
		for _, s := range proofs[idx] {
			p = append(p, common.HexToHash(s))
		}
		assert.True(t, VerifyProof(idx, balance.Account, balance.Amount, p, root))
	}

	info, err := ParseBalanceMap(balances)
	assert.Nil(t, err)
	assert.Equal(t, info.MerkleRoot.Hex(), "0x2ec9c2fc2a55df417ba88ecd833f165fa3c5941772ebaf8c5f4debe33f4d1b12")
	assert.Equal(t, info.TokenTotal, "0x2ee")
}
