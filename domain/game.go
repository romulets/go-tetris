package domain

type Tile rune

const (
	None   = Tile('n')
	Red    = Tile(zBlock)
	Orange = Tile(lBlock)
	Yellow = Tile(oBlock)
	Cyan   = Tile(iBlock)
	Green  = Tile(sBlock)
	Blue   = Tile(jBlock)
	Purple = Tile(tBlock)
)

type gameStatus byte

const (
	gameRunning       = 0
	gameOver          = 1
	inconsistentState = 2
)

type coord struct {
	x, y int
}

type movingBlock struct {
	block block
	coord coord
}

type board [20][10]Tile

type blockQueue [3]block

type Game struct {
	board        board
	currentBlock movingBlock
	heldBlock    *block
	blockQueue   blockQueue
	randomBlock  func() block
}

// Returns the first element of the queue and adds the argument to last position
func (q *blockQueue) pollAndAdd(lastBlock block) block {
	firstBlock := q[0]
	for x := 1; x < len(q); x++ {
		q[x-1] = q[x]
	}

	q[len(q)-1] = lastBlock

	return firstBlock
}

func newGame(randomBlock func() block) *Game {
	var board board
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			board[x][y] = None
		}
	}

	var queue blockQueue
	for x := 0; x < len(queue); x++ {
		queue[x] = randomBlock()
	}

	game := &Game{
		board:       board,
		blockQueue:  queue,
		randomBlock: randomBlock,
	}

	game.spawnBlock()

	return game
}

func (g *Game) spawnBlock() gameStatus {
	nextBlock := g.blockQueue.pollAndAdd(g.randomBlock())
	g.currentBlock = movingBlock{
		block: nextBlock,
		coord: initialPos(nextBlock),
	}

	// Same coordinates because it's the first render
	return g.processNewCoord(g.currentBlock.coord)
}

func initialPos(nextBlock block) coord {
	initialPos := coord{0, 3}
	if nextBlock.shape == oBlock {
		initialPos = coord{0, 4}
	}
	return initialPos
}

func (g *Game) processNewCoord(newCoord coord) gameStatus {
	currentCoord := g.currentBlock.coord

	// Clean up current coordinate
	for rowIdx, row := range g.currentBlock.block.body {
		for colIdx, col := range row {
			if col {
				// what if the game is inconsistent?
				//				if g.board[currentCoord.y+colIdx][newCoord.x+x] != Tile(g.currentBlock.block.shape) {
				//					return inconsistentState
				//				}
				g.board[currentCoord.y+colIdx][currentCoord.x+rowIdx] = None
			}
		}
	}

	// "Plot" new coordinate
	g.currentBlock.coord = newCoord
	for x, row := range trimBlock(g.currentBlock.block) {
		for y, col := range row {
			if col {
				if g.board[newCoord.x+x][newCoord.y+y] != None {
					return gameOver
				}
				g.board[newCoord.x+x][newCoord.y+y] = Tile(g.currentBlock.block.shape)
			}
		}
	}

	return gameRunning
}

func trimBlock(b block) [][]bool {
	// copy
	trimBody := make([][]bool, len(b.body))
	for x := range b.body {
		trimBody[x] = make([]bool, len(b.body[x]))
		copy(trimBody[x], b.body[x])
	}

	// remove top horizontals
	for _, row := range b.body {
		hasContent := false
		for _, cel := range row {
			hasContent = hasContent || cel
			if hasContent {
				break
			}
		}

		if !hasContent {
			trimBody = trimBody[1:]
		} else {
			break
		}
	}

	// remove bottom horizontals
	// remove left verticals
	// remove right verticals

	return trimBody
}
