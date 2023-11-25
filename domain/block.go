package domain

import (
	"crypto/rand"
	"math/big"
)

type shape rune

const (
	zBlock shape = 'z'
	lBlock shape = 'l'
	oBlock shape = 'o'
	iBlock shape = 'i'
	sBlock shape = 's'
	jBlock shape = 'j'
	tBlock shape = 't'
)

var allBuilders = []func() block{
	buildLBlock,
	buildJBlock,
	buildIBlock,
	buildOBlock,
	buildSBlock,
	buildZBlock,
	buildTBlock,
}

type block struct {
	shape shape
	body  [][]bool
}

func (b *block) rotate() {
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

func buildRandomBlock() block {
	i, _ := rand.Int(rand.Reader, big.NewInt(int64(len(allBuilders))))
	return allBuilders[i.Int64()]()
}

func buildIBlock() block {
	return block{
		shape: iBlock,
		body: [][]bool{
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
			{true, true, true, true},
		},
	}
}

func buildJBlock() block {
	return block{
		shape: jBlock,
		body: [][]bool{
			{false, false, false},
			{true, false, false},
			{true, true, true},
		},
	}
}

func buildLBlock() block {
	return block{
		shape: lBlock,
		body: [][]bool{
			{false, false, false},
			{false, false, true},
			{true, true, true},
		},
	}
}

func buildOBlock() block {
	return block{
		shape: oBlock,
		body: [][]bool{
			{true, true},
			{true, true},
		},
	}
}

func buildSBlock() block {
	return block{
		shape: sBlock,
		body: [][]bool{
			{false, false, false},
			{false, true, true},
			{true, true, false},
		},
	}
}

func buildZBlock() block {
	return block{
		shape: zBlock,
		body: [][]bool{
			{false, false, false},
			{true, true, false},
			{false, true, true},
		},
	}
}

func buildTBlock() block {
	return block{
		shape: tBlock,
		body: [][]bool{
			{true, true, true},
			{false, true, false},
			{false, true, false},
		},
	}
}
