package proxy

import (
	"errors"
	"sync"
	"log"
)

type MemoryBuffer struct {
	blockSize int
	usedBlocks []bool
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

	usedBlocks := make([]bool, bufferSize/blockSize)

	return &MemoryBuffer{
		usedBlocks: usedBlocks,
		blockSize: blockSize,
		bytes:     make([]byte, bufferSize),
		mux:       sync.Mutex{},
	}, nil
}

func (b *MemoryBuffer) GetNextBlock() *Block {
	b.mux.Lock()
	defer b.mux.Unlock()
	
	i := b.nextBlockIndex()
	bytes := b.bytes[i:b.blockSize]

	// log.Println("DEBUG: getting block", i)

	return &Block{
		index: i,
		Bytes: bytes,
	}
}

func (b *MemoryBuffer) ReturnBlock(block *Block) {
	if block == nil {
		return
	}

	b.mux.Lock()
	defer b.mux.Unlock()

	// log.Println("DEBUG: returning block", block.index)

	b.usedBlocks[block.index] = false
}

func (b *MemoryBuffer) nextBlockIndex() int {
	for n, x := range b.usedBlocks {
		if x == false {
			b.usedBlocks[n] = true
			return n
		}
	}

	log.Panic("buffer overflow!")
	return -1
}
