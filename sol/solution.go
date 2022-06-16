package sol

type Pair struct {
	row, col int
}

func solve(board [][]byte) {
	ROW := len(board)
	COL := len(board[0])
	lives := make(map[Pair]struct{})

	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row >= ROW || col < 0 || col >= COL || board[row][col] == 'X' {
			return
		}
		cell := Pair{row: row, col: col}
		if _, ok := lives[cell]; ok {
			return
		}
		lives[cell] = struct{}{}
		dfs(row-1, col)
		dfs(row+1, col)
		dfs(row, col-1)
		dfs(row, col+1)
	}
	for row := 0; row < ROW; row++ {
		dfs(row, 0)
		dfs(row, COL-1)
	}
	for col := 0; col < COL; col++ {
		dfs(0, col)
		dfs(ROW-1, col)
	}
	for row := 0; row < ROW; row++ {
		for col := 0; col < COL; col++ {
			if _, ok := lives[Pair{row: row, col: col}]; !ok && board[row][col] == 'O' {
				board[row][col] = 'X'
			}
		}
	}
}
