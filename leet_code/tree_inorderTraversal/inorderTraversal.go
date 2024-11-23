package main

import "fmt"

// TreeNode structure for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Main function to test the inorder traversal
func main() {
	// Construct the binary tree as per the example
	/*
	             1
	           /   \
	          2     3
	         / \
	        4   5
	*/
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	// Call the inorderTraversal function
	result := inorderTraversal(root)

	// Print the result
	fmt.Println("Inorder Traversal:", result)
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	stk := []*TreeNode{}
	cu := root

	for cu != nil || len(stk) > 0 {

		for cu != nil {
			stk =append(stk, cu)
			cu = cu.Left
		}
		cu = stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		res = append(res, cu.Val)

		cu = cu.Right

	}
	return res 
}