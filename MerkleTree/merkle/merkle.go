package merkle

import (
	"crypto/sha256"
	"fmt"
)

type Merklenode struct {
	Left  *Merklenode
	Right *Merklenode
	Data  []byte
}
type Merkletree struct {
	Rootnode *Merklenode
}

func NewMerklenode(left *Merklenode, right *Merklenode, data []byte) *Merklenode {
	nnode := Merklenode{}
	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		nnode.Data = hash[:]
	} else {
		prehash := append(left.Data, right.Data...) //zhuijia
		hash := sha256.Sum256(prehash)
		nnode.Data = hash[:]
	}
	nnode.Left = left
	nnode.Right = right
	return &nnode
}
func NewMerkletree(data [][]byte) *Merkletree {
	var nodes []Merklenode
	if len(data)%2 == 1 {
		data = append(data, data[len(data)-1])
	}
	for _, da := range data {
		node := NewMerklenode(nil, nil, da)
		nodes = append(nodes, *node)
	} //get all leave nodes
	for i := 0; i < len(data)/2; i++ {
		var newLevel []Merklenode
		for j := 0; j < len(nodes); j += 2 {
			node := NewMerklenode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}
		if len(newLevel)%2 == 1 {
			newLevel = append(newLevel, newLevel[len(newLevel)-1])
		} //if num is odd,copy the last one
		nodes = newLevel
	} //log or div2?
	ntree := Merkletree{&nodes[0]}
	return &ntree
}
func Printnode(node *Merklenode) {
	fmt.Printf("%p\n", node)
	if node != nil {
		fmt.Printf("left[%p],right[%p],data(%x)\n", node.Left, node.Right, node.Data)
	}
}
func ShowMerkletree(root *Merklenode) {
	if root == nil {
		return
	} else {
		fmt.Printf("%p\n", root)
		if root != nil {
			fmt.Printf("left[%p],right[%p],data(%x)\n", root.Left, root.Right, root.Data)
		}
	}
	ShowMerkletree(root.Left)
	ShowMerkletree(root.Right)
}
