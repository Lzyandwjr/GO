package main

/*
根据 百度百科 ，生命游戏，简称为生命，是英国数学家约翰·何顿·康威在 1970 年发明的细胞自动机。

给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。每个细胞都具有一个初始状态：1 即为活细胞（live），或 0 即为死细胞（dead）。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：

如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
根据当前状态，写一个函数来计算面板上所有细胞的下一个（一次更新后的）状态。下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是同时发生的。

 

示例：

输入：
[
  [0,1,0],
  [0,0,1],
  [1,1,1],
  [0,0,0]
]
输出：
[
  [0,0,0],
  [1,0,1],
  [0,1,1],
  [0,1,0]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/game-of-life
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func gameOfLife(board [][]int){
	for i := 0 ;i<len(board);i++{
		for j := 0 ;j < len(board[0]);j++{
			up := i-1 //上方
			left := j-1
			right := i+1
			down := j+1
			num := 0
			if up >=0 {
				if board[up][j] >= 1{
					num ++
				}
				if left >=0 && board[up][left]>=1{
					num ++
				}
				if down < len(board[0]) && board[up][down] >= 1{
					num ++
				}
			}
			if left >= 0 && board[i][left] >= 1{
				num ++
			}
			if down < len(board[0]) && board[i][down] >= 1{
				num ++
			}
			if right < len(board){
				if down <len(board[0]) && board[right][down] >=1 {
					num ++
				}
				if board[right][j] >= 1{
					num ++
				}
				if left >= 0 && board [right][left] >=1 {
					num ++
				}
			}
			if (num < 2 || num >3 ) && board[i][j] ==1 {
				board[i][j] = 2
			}else if (num ==2 || num ==3) && board[i][j] ==1 {
				board[i][j] = 1
			}else if num ==3 && board[i][j] == 0{
				board[i][j]	= -2
			}
		}
	}
	for i :=0;i<len(board);i++{
		for j := 0;j<len(board[0]);j++{
			if board[i][j] == 2{
				board[i][j] =0
			}else if board[i][j] == -2{
				board[i][j] =1
			}
		}
	}
}

func main(){

}