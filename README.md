# merkle-distributor
This package is the Go equivalent of [github.com/Uniswap/merkle-distributor](https://github.com/Uniswap/merkle-distributor).

## Install
```bash
go get -u github.com/fachebot/merkle-distributor
```

## Getting started

Simple example:
```go
package main

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	distributor "github.com/fachebot/merkle-distributor"
)

func main() {
	balances := []distributor.Balance{
		{Account: common.HexToAddress("0x17ec8597ff92C3F44523bDc65BF0f1bE632917ff"), Amount: big.NewInt(200)},
		{Account: common.HexToAddress("0x63FC2aD3d021a4D7e64323529a55a9442C444dA0"), Amount: big.NewInt(300)},
		{Account: common.HexToAddress("0xD1D84F0e28D6fedF03c73151f98dF95139700aa7"), Amount: big.NewInt(250)},
	}
	info, err := distributor.ParseBalanceMap(balances)
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
```

Output:
```json
{
	"merkleRoot": "0x2ec9c2fc2a55df417ba88ecd833f165fa3c5941772ebaf8c5f4debe33f4d1b12",
	"tokenTotal": "0x2ee",
	"claims": [{
		"index": 0,
		"amount": "0xc8",
		"proof": ["0x2a411ed78501edb696adca9e41e78d8256b61cfac45612fa0434d7cf87d916c6"]
	}, {
		"index": 1,
		"amount": "0x12c",
		"proof": ["0xbfeb956a3b705056020a3b64c540bff700c0f6c96c55c0a5fcab57124cb36f7b", "0xd31de46890d4a77baeebddbd77bf73b5c626397b73ee8c69b51efe4c9a5a72fa"]
	}, {
		"index": 2,
		"amount": "0xfa",
		"proof": ["0xceaacce7533111e902cc548e961d77b23a4d8cd073c6b68ccf55c62bd47fc36b", "0xd31de46890d4a77baeebddbd77bf73b5c626397b73ee8c69b51efe4c9a5a72fa"]
	}]
}
```

## License

MIT
