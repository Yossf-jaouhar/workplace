package main

import "fmt"

// TreeNode structure for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Main function to build the tree and test the traversal
func main() {
	// Construct the binary tree as per the input
	/*
	             1
	           /   \
	          2     3
	         / \      \
	        4   5      8
	           / \    /
	          6   7  9
	*/
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Left.Right.Left = &TreeNode{Val: 6}
	root.Left.Right.Right = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 8}
	root.Right.Right.Left = &TreeNode{Val: 9}

	// Call the preorder traversal function
	result := postorderTraversal(root)

	// Print the result
	fmt.Println("Preorder Traversal:", result)
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
    if root == nil {
        return res
    }

    st := []*TreeNode{root}

    for len(st)  > 0 {
        
        node := st[len(st)-1]
        st = st[:len(st)-1]

        res = append(res , node.Val)
        if node.Left != nil  {
            st = append(st, node.Left)
        }

		if node.Right != nil  {
            st = append(st, node.Right)
        }

        
    }
    return res
}


func postorderTraversal(root *TreeNode) []int {
    var res []int

    if root == nil {
        return res
    }

    stk := []*TreeNode{root}

    for len(stk) > 0 {
        node := stk[len(stk)-1]
        stk = stk[:len(stk)-1]
  
        if node != nil {
            stk = append(stk, node.Left)
            stk = append(stk, node.Right)
			res = append([]int{node.Val}, res...)
        }
    }
   
    return res
}
