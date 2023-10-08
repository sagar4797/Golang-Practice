package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}

	return dfs(root, head) || isSubPath(head.Next, root.Left) || isSubPath(head.Next, root.Right)
}

func dfs(root *TreeNode, node *ListNode) bool {
	if node == nil {
		return true
	}

	if root == nil {
		return false
	}

	if node.Val == root.Val {
		return dfs(root.Left, node.Next) || dfs(root.Right, node.Next)
	}

	return false
}

func main() {
	list := &ListNode{4, &ListNode{2, &ListNode{8, nil}}}

	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	isExist := isSubPath(list, root1)
	fmt.Printf("\nIs Exists: %+v", isExist)
}
