package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Constructing a binary tree:
	//         3
	//        / \
	//       9   20
	//          /  \
	//         15   7
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	// Calling the function
	result := averageOfLevels(root)

	// Printing the result
	fmt.Println("Average of levels:", result)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	
	var res []int
	var result []float64
	if root == nil {
		return result
	}

	nodesize := []*TreeNode{root}

	for len(nodesize) > 0 {

		size := len(nodesize)

		for i := 0; i < size; i++ {
			node := nodesize[0]
			nodesize = nodesize[1:]
			res = append(res, node.Val)

			if node.Left != nil {
				nodesize = append(nodesize, node.Left)
			}

			if node.Right != nil {
				nodesize = append(nodesize, node.Right)
			}

		}
		result = append(result, hhh(res))
		res = nil
	}

	fmt.Println(result)
	return result
}

func hhh(res []int) float64 {
	v := 0

	for _, char := range res {
		v += char
	}
	b := float64(v) / float64(len(res))

	return b
}	
