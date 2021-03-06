package blockchain

import (
	"bytes"
	"encoding/gob"
	"github.com/wdnb/gene/gecko"
	"log"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Timestamp     int64
	Gecko []*gecko.Gecko
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	Height        int
}

// NewBlock creates and returns Block
func NewBlock(transactions []*Transaction,gecko []*gecko.Gecko, prevBlockHash []byte, height int) *Block {

	block := &Block{time.Now().Unix(),gecko, transactions,prevBlockHash, []byte{}, 0, height}
	//禁用pow
	//fmt.Println(block)
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	//os.Exit(22222222)
	block.Hash = hash[:]
	block.Nonce = nonce

	//var hash [32]byte

	//fmt.Println(string(prevBlockHash))
	//block.Hash=sha256.Sum256([32]byte(prevBlockHash))

	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase},[]*gecko.Gecko{{nil,0,"", struct{}{}, struct{}{},"神之壁虎","",0,0,nil}}, []byte{}, 0)
}

// HashTransactions returns a hash of the transactions in the block
func (b *Block) HashTransactions() []byte {
	var transactions [][]byte

	for _, tx := range b.Transactions {
		transactions = append(transactions, tx.Serialize())
	}
	mTree := NewMerkleTree(transactions)

	return mTree.RootNode.Data
}

// Serialize serializes the block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// DeserializeBlock deserializes a block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
