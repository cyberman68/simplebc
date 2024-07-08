package blockchain

import (
	"go-demo/blockchain/core"
	"strconv"
)

func main(){
    	chain := core.NewBlockChain()
	for i := 0; i < 10; i++ {
		chain.SendData("block:" + strconv.Itoa(i))
	}
}
