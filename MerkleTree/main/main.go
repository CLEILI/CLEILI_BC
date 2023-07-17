package main

import (
	"CLEILI_BC/MerkleTree/merkle"
	"fmt"
)

func main() {
	data := [][]byte{
		[]byte("node1"),
		[]byte("node2"),
		[]byte("node3"),
		[]byte("node4"),
		[]byte("node5"),
		[]byte("node6"),
	}
	tree := merkle.NewMerkletree(data)
	merkle.ShowMerkletree(tree.Rootnode)
	fmt.Printf("done")
}
