package proxy

import (
	"errors"
	"sync"
	"time"
)

type MemoryBuffer struct {
	nextBlock int
	blockSize int
	bytes     []byte
	mux       sync.Mutex
}

type Block struct {
	index int
	Bytes []byte
}

func NewMemoryBuffer(bufferSize, blockSize int) (*MemoryBuffer, error) {
	if blockSize > bufferSize {
		return nil, errors.New("buffersize must be larger than or equal to blocksize")
	}

	return &MemoryBuffer{
		blockSize: blockSize,
		bytes:     make([]byte, bufferSize),
		mux:       sync.Mutex{},
	}, nil
}

func (b *MemoryBuffer) GetNextBlock() *Block {
	for b.nextBlock-1 >= len(b.bytes)/b.blockSize {
		// wait for a memory block to become available
		time.Sleep(time.Millisecond * 5)
	}

	b.mux.Lock()
	defer b.mux.Unlock()

	bytes := b.bytes[b.nextBlock:b.blockSize]

	block := &Block{
		index: b.nextBlock,
		Bytes: bytes,
	}

	b.nextBlock++

	return block
}

func (b *MemoryBuffer) ReturnBlock(block *Block) {
	if block == nil {
		return
	}

	b.mux.Lock()
	defer b.mux.Unlock()

	b.nextBlock = block.index
}
