package proxy

import (
	"errors"
	"sync"
)

// Pool of bytes that requests can use to buffer responses
type Pool struct {
	blockSize  int
	usedBlocks []bool
	bytes      []byte
	mux        sync.Mutex
}

// Block of the pool that can be used to buffer a response
type Block struct {
	index int
	Bytes []byte
}

// NewPool creates a new poroperly initialised pool
func NewPool(bufferSize, blockSize int) (*Pool, error) {
	if blockSize > bufferSize {
		return nil, errors.New("buffersize must be larger than or equal to blocksize")
	}

	usedBlocks := make([]bool, bufferSize/blockSize)

	return &Pool{
		usedBlocks: usedBlocks,
		blockSize:  blockSize,
		bytes:      make([]byte, bufferSize),
		mux:        sync.Mutex{},
	}, nil
}

// GetNextBlock from the pool
func (b *Pool) GetNextBlock() *Block {
	b.mux.Lock()
	defer b.mux.Unlock()

	i := b.nextBlockIndex()

	return &Block{
		index: i,
		Bytes: b.bytes[i:b.blockSize],
	}
}

// ReturnBlock to the pool
func (b *Pool) ReturnBlock(block *Block) {
	if block == nil {
		return
	}

	b.mux.Lock()
	defer b.mux.Unlock()

	// log.Println("DEBUG: returning block", block.index)

	b.usedBlocks[block.index] = false
}

func (b *Pool) nextBlockIndex() int {
	for n, x := range b.usedBlocks {
		if x == false {
			b.usedBlocks[n] = true
			return n
		}
	}

	return -1
}
