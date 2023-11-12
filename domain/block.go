package domain

import (
	"crypto/rand"
	"math/big"
)

type shape rune

const (
	lBlock shape = 'l'
	jBlock shape = 'j'
	iBlock shape = 'i'
	oBlock shape = 'o'
	sBlock shape = 's'
	zBlock shape = 'z'
	tBlock shape = 't'
)

var allBuilders = []func() Block{
	buildLBlock,
	buildJBlock,
	buildIBlock,
	buildOBlock,
	buildSBlock,
	buildZBlock,
	buildTBlock,
}

type Block struct {
	shape shape
	body  [][]bool
}

func (b *Block) Rotate() {
	n := len(b.body)

	for x := 0; x < n/2; x++ {
		for y := x; y < n-x-1; y++ {
			temp := b.body[x][y]

			b.body[x][y] = b.body[y][n-1-x]
			b.body[y][n-1-x] = b.body[n-1-x][n-1-y]
			b.body[n-1-x][n-1-y] = b.body[n-1-y][x]
			b.body[n-1-y][x] = temp
		}
	}
}

func BuildRandomBlock() Block {
	i, _ := rand.Int(rand.Reader, big.NewInt(int64(len(allBuilders))))
	return allBuilders[i.Int64()]()
}

func buildIBlock() Block {
	return Block{
		shape: iBlock,
		body: [][]bool{
			{true, false, false, false},
			{true, false, false, false},
			{true, false, false, false},
			{true, false, false, false},
		},
	}
}

func buildJBlock() Block {
	return Block{
		shape: jBlock,
		body: [][]bool{
			{false, false, false, false},
			{false, false, false, false},
			{true, false, false, false},
			{true, true, true, true},
		},
	}
}

func buildLBlock() Block {
	return Block{
		shape: lBlock,
		body: [][]bool{
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, true},
			{true, true, true, true},
		},
	}
}

func buildOBlock() Block {
	return Block{
		shape: oBlock,
		body: [][]bool{
			{true, true},
			{true, true},
		},
	}
}

func buildSBlock() Block {
	return Block{
		shape: sBlock,
		body: [][]bool{
			{false, true, true},
			{false, true, false},
			{true, true, false},
		},
	}
}

func buildZBlock() Block {
	return Block{
		shape: zBlock,
		body: [][]bool{
			{true, true, false},
			{false, true, false},
			{false, true, true},
		},
	}
}

func buildTBlock() Block {
	return Block{
		shape: tBlock,
		body: [][]bool{
			{true, true, true},
			{false, true, false},
			{false, true, false},
		},
	}
}
