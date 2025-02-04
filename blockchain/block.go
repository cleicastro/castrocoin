package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type Block struct {
	Hash    []byte
	Data    []byte
	PrevHas []byte
	Nonce   int
}

func CreateBlock(data string, prevHas []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHas, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHas}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic()
	}
}
