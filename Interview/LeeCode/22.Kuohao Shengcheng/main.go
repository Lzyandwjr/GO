package main
//深度优先遍历 从小开始
func generateParenthesis(n int)[]string{
	res := []string{}
	if n == 0{
		return res
	}
	DFS("",0,0,n,&res)
	return res
}

func DFS(s string,left,right,n int,res *[]string){
	if left == n && right == n{
		*res =append(*res,s)
		return
	}
	if left < right{
		return
	}
	if left < n {
		DFS(s+"(",left+1,right,n,res)
	}
	if right < n{
		DFS(s+")",left,right+1,n,res)
	}
}

// 动态规划
func generateParenthesisL(n int)[]string{
	maps := make([][]string,n+1)
	maps[0] = []string{""}
	for i := 1;i<=n;i++{
		maps[i] = []string{}
		for j:=0;j<i;j++{
			left,right := j,i-1-j
			for _,value := range maps[left]{
				for _,value2 := range maps[right]{
					maps[i] = append(maps[i],"("+value+")"+value2)
				}
			}
		}
	}
	return maps[n]

}
//深度优先遍历 从最大开始
func generateParenthesisR(n int) []string {
	res := []string{}
	if n == 0{
		return res
	}
	dfs("",n,n,&res)
	return res
}
func dfs(ss string,left,right int, res *[]string)  {
	if left == 0 && right == 0{
		*res = append(*res,ss)
		return
	}
	if left  > right{
		return
	}
	if left > 0{
		dfs(ss +"(",left-1,right,res)
	}
	if right > 0{
		dfs(ss +")",left,right-1,res)
	}
}

func main(){

}
