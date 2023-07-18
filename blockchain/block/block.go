package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64
	Data      []byte
	Prehash   []byte
	Hash      []byte
	Nonce     int64
}

func (b *Block) SetHash() {
	ts := []byte(strconv.FormatInt(b.Timestamp, 10))
	header := bytes.Join([][]byte{b.Data, b.Prehash, b.Hash, ts}, []byte{})
	hash := sha256.Sum256(header)
	b.Hash = hash[:]
}
func Newblock(data string, prehash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prehash, []byte{}, 0}
	thepow := Newproofofwork(block)
	nonce, hash := thepow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return block
}
func Newgenesisblock() *Block {
	return Newblock("genesis block", []byte{})
}

//next is pow
var maxNonce = math.MaxInt64

const targetBits = 16

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func Newproofofwork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{block: b, target: target}
	return pow
} //get a standard of pow
func (pow *ProofOfWork) Run() (int64, []byte) {
	nonce := 0
	var hashint big.Int //big data compare
	var hash [32]byte   //why 32
	fmt.Printf("the block containing %s,maxnonce is %d\n", pow.block.Data, maxNonce)
	for nonce < maxNonce {
		ts := []byte(strconv.FormatInt(pow.block.Timestamp, 10))
		tb := []byte(strconv.FormatInt(int64(targetBits), 10))
		non := []byte(strconv.FormatInt(int64(nonce), 10))
		data := bytes.Join([][]byte{pow.block.Data, pow.block.Prehash, ts, tb, non}, []byte{})
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashint.SetBytes(hash[:])
		if hashint.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	fmt.Println()
	return int64(nonce), hash[:]
}
func (pow *ProofOfWork) Validate() bool {
	var hashint big.Int
	ts := []byte(strconv.FormatInt(pow.block.Timestamp, 10))
	tb := []byte(strconv.FormatInt(int64(targetBits), 10))
	non := []byte(strconv.FormatInt(int64(pow.block.Nonce), 10))
	data := bytes.Join([][]byte{pow.block.Data, pow.block.Prehash, ts, tb, non}, []byte{})
	hash := sha256.Sum256(data)
	hashint.SetBytes(hash[:])
	return hashint.Cmp(pow.target) == -1
}
