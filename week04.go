//被围绕的区域（Medium)
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return      
	}
	r, c := len(board), len(board[0]) 
	for i := 0; i < r; i++ {      
		dfs(board, r, c, i, 0)    
		dfs(board, r, c, i, c-1)  
	}
	for i := 1; i < c-1; i++ {     
		dfs(board, r, c, 0, i)    
		dfs(board, r, c, r-1, i)  
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {   
			switch board[i][j] {
			case 'A':
				board[i][j] = 'O'   
			case 'O':
				board[i][j] = 'X'   
			}
		}
	}
}

func dfs(board [][]byte, r, c, x, y int) {
	if x < 0 || x >= r || y < 0 || y >= c || board[x][y] != 'O' {
		return  
	}
	board[x][y] = 'A'  
	dfs(board, r, c, x-1, y)
	dfs(board, r, c, x+1, y)
	dfs(board, r, c, x, y-1) 
	dfs(board, r, c, x, y+1)
}

//把二叉搜索树转换为累加树（Medium）
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func convertBST(root *TreeNode) *TreeNode {
    sum := 0
    return convertTree(root, &sum)
}

func convertTree(root *TreeNode, sum *int) *TreeNode {
    if root == nil {
        return nil
    }
    convertTree(root.Right, sum)
    *sum = *sum + root.Val
    root.Val = *sum
    convertTree(root.Left, sum)
    return root
}
