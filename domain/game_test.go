package domain

import (
	"reflect"
	"strings"
	"testing"
)

func Test_newGame_OFirst(t *testing.T) {
	g := newGame(func() block {
		return buildOBlock()
	})

	if g.currentBlock.block.shape != oBlock {
		t.Errorf("expected shape %c but got %c", oBlock, g.currentBlock.block.shape)
	}

	if !reflect.DeepEqual(g.currentBlock.coord, coord{x: 4, y: 0}) {
		t.Errorf(
			"expected coord (%2d,%2d) but got (%2d,%2d)",
			4,
			0,
			g.currentBlock.coord.x,
			g.currentBlock.coord.y,
		)
	}

	assertBoardIs(t, g, [20][10]Tile{
		{None, None, None, None, Yellow, Yellow, None, None, None, None},
		{None, None, None, None, Yellow, Yellow, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
	})
}

func Test_newGame_JFirst(t *testing.T) {
	t.Skip("Not implemented yet")

	g := newGame(func() block {
		return buildJBlock()
	})

	if g.currentBlock.block.shape != oBlock {
		t.Errorf("expected shape %c but got %c", oBlock, g.currentBlock.block.shape)
	}

	if !reflect.DeepEqual(g.currentBlock.coord, coord{x: 4, y: 0}) {
		t.Errorf(
			"expected coord (%2d,%2d) but got (%2d,%2d)",
			4,
			0,
			g.currentBlock.coord.x,
			g.currentBlock.coord.y,
		)
	}

	assertBoardIs(t, g, [20][10]Tile{
		{None, None, None, Blue, None, None, None, None, None, None},
		{None, None, None, Blue, Blue, Blue, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
		{None, None, None, None, None, None, None, None, None, None},
	})
}

func assertBoardIs(t *testing.T, g *Game, board [20][10]Tile) {
	t.Helper()

	diff := false
	for colIdx, col := range g.board {
		for rowIdx, tile := range col {
			if board[colIdx][rowIdx] != tile {
				diff = true
				t.Errorf("Expected tile %c but got %c (coord{%2d,%2d})", board[colIdx][rowIdx], tile, colIdx, rowIdx)
			}
		}
	}

	if diff {
		t.Errorf("Boards doesn't match \n\nEXPECTED\n%s, \nGOT\n%s", fmtBoard(board), fmtBoard(g.board))
	}
}

func fmtBoard(board [20][10]Tile) string {
	b := strings.Builder{}
	for _, row := range board {
		for _, tile := range row {
			if tile == None {
				b.WriteRune('.')
			} else {
				b.WriteRune(rune(tile))
			}
		}
		b.WriteRune('\n')
	}

	return b.String()
}
