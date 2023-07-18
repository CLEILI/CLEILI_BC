package main

import (
	"CLEILI_BC/blockchain/block"
	"fmt"
)

type Blockchain struct {
	blocks []*block.Block
}

func (bc *Blockchain) Addblock(data string) {
	preblockp := bc.blocks[len(bc.blocks)-1]
	newblock := block.Newblock(data, preblockp.Hash)
	bc.blocks = append(bc.blocks, newblock)
}
func Newblockchain() *Blockchain {
	return &Blockchain{[]*block.Block{block.Newgenesisblock()}}
}
func main() {
	bc := Newblockchain()
	bc.Addblock("send 100 yuan to lei")
	bc.Addblock("send 10 yuan to ze")
	for _, block := range bc.blocks {
		fmt.Printf("time:%v\n", block.Timestamp)
		fmt.Printf("prehash:%x\n", block.Prehash)
		fmt.Printf("data:%s\n", block.Data)
		fmt.Printf("Hash:%p\n", block.Hash)
	}
}
