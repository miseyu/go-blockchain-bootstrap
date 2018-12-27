package lib

import "time"

type Block struct {
	Data         string
	Hash         string
	previousHash string
	timestamp    int64
}

func New(data string, previousHash string) *Block {
	return &Block{
		Data:         data,
		previousHash: previousHash,
		timestamp:    time.Now().UnixNano() / int64(time.Millisecond),
	}
}
