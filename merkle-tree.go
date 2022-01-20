package distributor

import (
	"bytes"
	"errors"
	"math"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func combinedHash(first, second common.Hash) (hash common.Hash) {
	var buffer [64]byte
	if bytes.Compare(first[:], second[:]) <= 0 {
		copy(buffer[:32], first[:])
		copy(buffer[32:], second[:])
		copy(hash[:], crypto.Keccak256(buffer[:]))
	} else {
		copy(buffer[:32], second[:])
		copy(buffer[32:], first[:])
		copy(hash[:], crypto.Keccak256(buffer[:]))
	}
	return hash
}

type MerkleTree struct {
	layers                     [][]common.Hash
	elements                   Elements
	bufferElementPositionIndex map[common.Hash]int
}

func NewMerkleTree(elements Elements) (*MerkleTree, error) {
	sort.Sort(elements)
	elements = elements.Dedup()

	bufferElementPositionIndex := make(map[common.Hash]int)
	for idx, el := range elements {
		bufferElementPositionIndex[el] = idx
	}

	var err error
	tree := MerkleTree{
		elements:                   elements,
		bufferElementPositionIndex: bufferElementPositionIndex,
	}
	tree.layers, err = tree.GetLayers(elements)
	if err != nil {
		return nil, err
	}

	return &tree, nil
}

func (m *MerkleTree) GetRoot() common.Hash {
	return m.layers[len(m.layers)-1][0]
}

func (m *MerkleTree) GetProof(el common.Hash) ([]common.Hash, error) {
	idx, ok := m.bufferElementPositionIndex[el]
	if !ok {
		return nil, errors.New("element does not exist in Merkle tree")
	}

	proof := make(Elements, 0)
	for _, layer := range m.layers {
		pairElement, ok := m.getPairElement(idx, layer)
		if ok {
			proof = append(proof, pairElement)
		}
		idx = int(math.Floor(float64(idx) / 2))
	}

	return proof, nil
}

func (m *MerkleTree) GetLayers(elements Elements) ([][]common.Hash, error) {
	if len(elements) == 0 {
		return nil, errors.New("empty tree")
	}

	layers := make([][]common.Hash, 0)
	layers = append(layers, elements)

	for len(layers[len(layers)-1]) > 1 {
		layers = append(layers, m.GetNextLayer(layers[len(layers)-1]))
	}
	return layers, nil
}

func (m *MerkleTree) GetNextLayer(elements Elements) Elements {
	layer := make(Elements, 0)
	for idx, el := range elements {
		if idx%2 == 0 {
			if idx+1 >= len(elements) {
				layer = append(layer, el)
			} else {
				layer = append(layer, combinedHash(el, elements[idx+1]))
			}
		}
	}
	return layer
}

func (m *MerkleTree) getPairElement(idx int, layer Elements) (common.Hash, bool) {
	pairIdx := idx - 1
	if idx%2 == 0 {
		pairIdx = idx + 1
	}

	if pairIdx >= len(layer) {
		return common.Hash{}, false
	}

	return layer[pairIdx], true
}
