package main

/**
 * Definition for a binary tree node. */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 && len(preorder) != len(inorder) {
		return nil
	}
	var rootplace int = preorder[0]
	var pos int
	for i := range inorder {
		if inorder[i] == rootplace {
			pos = i
			break
		}
	}
	root := &TreeNode{Val: rootplace}
	root.Left = buildTree(preorder[1:pos+1], inorder[:pos])
	root.Right = buildTree(preorder[pos+1:], inorder[pos+1:])
	return root

}
