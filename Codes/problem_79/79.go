package leetcode

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	// 二维数组，用来标记元素是否找过
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// 遍历二维数组
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, visited, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, visited [][]bool, word string, curr, i, j int) bool {
	// 没有找过当前元素 && 当前元素 == 字符串对应位置
	if !visited[i][j] && board[i][j] == word[curr] {
		// 到word最后一个字母了
		if curr == len(word)-1 {
			return true
		}

		// 当前元素设为访问过了
		visited[i][j] = true

		// Top
		if i-1 >= 0 && dfs(board, visited, word, curr+1, i-1, j) {
			return true
		}

		// Down
		if i+1 < len(board) && dfs(board, visited, word, curr+1, i+1, j) {
			return true
		}

		// Left
		if j-1 >= 0 && dfs(board, visited, word, curr+1, i, j-1) {
			return true
		}

		// Right
		if j+1 < len(board[0]) && dfs(board, visited, word, curr+1, i, j+1) {
			return true
		}

		// 重置
		visited[i][j] = false
	}
	return false
}
