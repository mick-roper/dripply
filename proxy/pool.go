package proxy

import (
	"sync"
)

const (
	initialPoolSize = 1
	blockSize       = 64 * 1024
)

// Pool of bytes that requests can use to buffer responses
type Pool struct {
	usedBlocks []bool
	blocks     []Block
	mux        sync.Mutex
}

// Block of the pool that can be used to buffer a response
type Block struct {
	index int
	Bytes []byte
}

// NewPool creates a new poroperly initialised pool
func NewPool(blockSize, maxBlocks int) (*Pool, error) {
	blocks := []Block{
		Block{Bytes: make([]byte, blockSize)},
	}
	usedBlocks := []bool{false}

	return &Pool{
		usedBlocks: usedBlocks,
		blocks:     blocks,
		mux:        sync.Mutex{},
	}, nil
}

// GetNextBlock from the pool
func (p *Pool) GetNextBlock() *Block {
	p.mux.Lock()
	defer p.mux.Unlock()

	var i int
	for i = p.nextBlockIndex(); i < 0; {
		p.grow()
	}

	p.usedBlocks[i] = true
	return &p.blocks[i]
}

// ReturnBlock to the pool
func (p *Pool) ReturnBlock(block *Block) {
	if block == nil {
		return
	}

	p.mux.Lock()
	defer p.mux.Unlock()

	// log.Println("DEBUG: returning block", block.index)

	p.usedBlocks[block.index] = false
}

func (p *Pool) nextBlockIndex() int {
	for n, x := range p.usedBlocks {
		if x == false {
			return n
		}
	}

	return -1
}

func (p *Pool) grow() {
	p.mux.Lock()
	defer p.mux.Unlock()

	l := len(p.blocks)

	// add more blocks
	p.blocks = append(p.blocks, make([]Block, l)...)
	for i := l - 1; i < len(p.blocks); i++ {
		p.blocks[i] = Block{Bytes: make([]byte, blockSize)}
	}

	// add more monitors
	p.usedBlocks = append(p.usedBlocks, make([]bool, l)...)
}
