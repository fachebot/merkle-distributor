package distributor

import (
	"bytes"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
)

type Elements []common.Hash

func (x Elements) Len() int           { return len(x) }
func (x Elements) Less(i, j int) bool { return bytes.Compare(x[i][:], x[j][:]) == -1 }
func (x Elements) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func (x Elements) Dedup() Elements {
	r := make(Elements, 0, len(x))
	for idx, el := range x {
		if idx == 0 || !bytes.Equal(x[idx-1][:], el[:]) {
			r = append(r, el)
		}
	}
	return r
}

func (x Elements) ToHexArray() []string {
	r := make([]string, 0, len(x))
	for _, el := range x {
		r = append(r, hex.EncodeToString(el[:]))
	}
	return r
}
