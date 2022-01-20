package distributor

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func ToNode(index int, account common.Address, amount *big.Int) common.Hash {
	var h common.Hash
	sha3 := solsha3.SoliditySHA3(
		[]string{"uint256", "address", "uint256"},
		[]interface{}{big.NewInt(int64(index)), account, amount},
	)
	copy(h[:], sha3)
	return h
}

func VerifyProof(index int, account common.Address, amount *big.Int, proof Elements, root common.Hash) bool {
	pair := ToNode(index, account, amount)
	for _, item := range proof {
		pair = combinedHash(pair, item)
	}

	return pair == root
}

type Balance struct {
	Account common.Address
	Amount  *big.Int
}

type BalanceTree struct {
	tree *MerkleTree
}

func NewBalanceTree(balances []Balance) (*BalanceTree, error) {
	elements := make(Elements, 0, len(balances))
	for idx, balance := range balances {
		elements = append(elements, ToNode(idx, balance.Account, balance.Amount))
	}

	tree, err := NewMerkleTree(elements)
	if err != nil {
		return nil, err
	}

	return &BalanceTree{tree: tree}, nil
}

func (b *BalanceTree) GetRoot() common.Hash {
	return b.tree.GetRoot()
}

func (b *BalanceTree) GetProof(index int, account common.Address, amount *big.Int) ([]common.Hash, error) {
	return b.tree.GetProof(ToNode(index, account, amount))
}
