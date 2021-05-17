package binarytree

import "fmt"

// TreeNode 树节点
type TreeNode struct {
	V     int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的遍历
// 种类： 前序，中序，后序三种，主要是针对的根节点的遍历顺序来进行区分；
// 左，中，右分别采用L,D,R来表示时，就是DLR, LDR, LRD

// 前序遍历

// DLRRecursion 前序遍历-递归
func DLRRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	// 先根节点
	fmt.Println(root.V)
	// 然后左子节点
	DLRRecursion(root.Left)
	// 然后右子节点
	DLRRecursion(root.Right)
}

// DLR 前序遍历
func DLR(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)             // 保存二叉树遍历的结果
	stack := make([]*TreeNode, 0)        // 保存子树的根节点，用于访问右节点
	for root != nil || len(stack) != 0 { // len(stack) 中如果不等于0 ，说明还有子树的根节点没出栈
		for root != nil {
			// 输出 根节点, 这里也可以先保存到切片
			result = append(result, root.V) // [root.v, root_l1.V, root_l2.V。。。]
			// 由于左节点遍历完成后，需要遍历遍历右边节点，所以需要一个切片来保存当前节点
			stack = append(stack, root)
			root = root.Left
		}
		// 从 stack 取出 保存节点，并向其右节点遍历
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

// LDR 中序遍历
func LDR(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	result := make([]int, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		result = append(result, node.V)
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}

// LRD 后序遍历
func LRD(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	var lastVisit *TreeNode

	for root != nil || len(stack) != 0 {

		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]

		if node.Right == nil || lastVisit == node.Right { // 说明子树的右节点为空或者已经遍历过了
			stack = stack[:len(stack)-1] // pop
			result = append(result, node.V)

			lastVisit = node // 标记当前节点已经弹出过
		} else { // 当前节点的右子节点存在，且没有被遍历过
			root = node.Right
		}
	}

	return nil
}
