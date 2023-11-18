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

var initialCoord = map[shape]func() coord{
	zBlock: func() coord { return coord{x: 0, y: 0} },
	lBlock: func() coord { return coord{x: 0, y: 0} },
	oBlock: func() coord { return coord{x: 4, y: 0} },
	iBlock: func() coord { return coord{x: 0, y: 0} },
	sBlock: func() coord { return coord{x: 0, y: 0} },
	jBlock: func() coord { return coord{x: 0, y: 0} },
	tBlock: func() coord { return coord{x: 0, y: 0} },
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
	current := g.blockQueue.pollAndAdd(g.randomBlock())
	g.currentBlock = movingBlock{
		block: current,
		coord: initialCoord[current.shape](),
	}

	// Same coordinates because it's the first render
	return g.processNewCoord(g.currentBlock.coord)
}

func (g *Game) processNewCoord(newCoord coord) gameStatus {
	currentCoord := g.currentBlock.coord

	// Clean up current coordinate
	for rowIdx, row := range g.currentBlock.block.body {
		for colIdx, col := range row {
			if col {
				// todo: what if the game is inconsistent?
				//				if g.board[currentCoord.y+colIdx][newCoord.x+rowIdx] != Tile(g.currentBlock.block.shape) {
				//					return inconsistentState
				//				}
				g.board[currentCoord.y+colIdx][currentCoord.x+rowIdx] = None
			}
		}
	}

	// "Plot" new coordinate
	g.currentBlock.coord = newCoord
	for rowIdx, row := range g.currentBlock.block.body {
		for colIdx, col := range row {
			if col {
				if g.board[newCoord.y+colIdx][newCoord.x+rowIdx] != None {
					return gameOver
				}
				g.board[newCoord.y+colIdx][newCoord.x+rowIdx] = Tile(g.currentBlock.block.shape)
			}
		}
	}

	return gameRunning
}
