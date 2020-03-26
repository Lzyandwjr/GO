package main



/*
	在 N * N 的网格上，我们放置一些 1 * 1 * 1  的立方体。

每个值 v = grid[i][j] 表示 v 个正方体叠放在对应单元格 (i, j) 上。

请你返回最终形体的表面积。

 

示例 1：

输入：[[2]]
输出：10
示例 2：

输入：[[1,2],[3,4]]
输出：34
示例 3：

输入：[[1,0],[0,2]]
输出：16
示例 4：

输入：[[1,1,1],[1,0,1],[1,1,1]]
输出：32
示例 5：

输入：[[2,2,2],[2,1,2],[2,2,2]]
输出：46

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/surface-area-of-3d-shapes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func surfaceArea(grid [][]int)int{
	N:=len(grid)
	ans:=0
	var max func(a int,b int)int
	max = func(a int, b int) int{
		if a>b{return a}
		return b
	}
	for i:=0;i<N;i++{
		for j:=0;j<N;j++{
			if grid[i][j]!=0{
				ans+=2
			}
			if i>0{
				ans+=max(grid[i][j]-grid[i-1][j],0)
			}else{
				ans+=grid[i][j]
			}
			if i<N-1{
				ans+=max(grid[i][j]-grid[i+1][j],0)
			}else{
				ans+=grid[i][j]
			}
			if j>0{
				ans+=max(grid[i][j]-grid[i][j-1],0)
			}else{
				ans+=grid[i][j]
			}
			if j<N-1{
				ans+=max(grid[i][j]-grid[i][j+1],0)
			}else{
				ans+=grid[i][j]
			}
		}
	}
	return ans
}

func main(){

}