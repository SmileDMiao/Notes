// 79. Word Search
// question
// 给一个二维数组(矩阵)和一个string，判断这个矩阵中是否包含这个string(矩阵中的元素连续起来出现组成string)
// example
// Input: [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], "ABCCED"; Output: true

// 思路
// 遍历二维数组，DFS(深度优先搜索)，查看每个元素可达的相邻元素

package main

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
func main() {
	board := [][]byte{[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")}
	word := "ABCCED"

	exist(board, word)
}
