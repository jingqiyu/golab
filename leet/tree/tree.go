package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var (
		result []int
	)
	if root == nil {
		return result
	} else {

		result = append(result, inorderTraversal(root.Left)...)
		result = append(result, root.Val)
		result = append(result, inorderTraversal(root.Right)...)
	}
	return result
}

/**
 *
			   3
			  / \
			 9  20
			   /  \
			  15   7
 *

	preorder = [3,9,20,15,7]
    inorder = [9,3,15,20,7]
 * param: []int preorder
 * param: []int inorder
 * return: *TreeNode
*/
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}

	lPre := len(preorder)

	// 1. 把pre顺序的元素入栈 先序的元素 第一个一定是一个树的根节点， 再通过看 中序遍历的第一个元素是否是相同节点
	// 如果节点不相同 证明先序当前元素后面的元素位于这个元素的左子树中
	// 如果节点相同 则要找到这个元素到底在刚才入栈的几个节点的哪个节点的右子树中

	stackArr := []*TreeNode{}
	stackArr = append(stackArr, root)

	var iOrderIdex = 0
	for i := 1; i < lPre; i++ {

		pVal := preorder[i]
		// peed操作 取栈顶的元素 是否和中序的元素相等
		node := stackArr[len(stackArr)-1]
		if node.Val != inorder[iOrderIdex] {
			node.Left = &TreeNode{Val: pVal}
			// push 操作 把刚刚处理好的元素放入到栈中
			stackArr = append(stackArr, node.Left)
		} else {
			//这里比较麻烦 ，之前已经入栈了几个元素， 我们要找到这个元素是位于哪个元素的右子树
			// 如何判断：
		}

	}

	return nil
}

// 递归比较好理解
// pI = 前序中的某个元素在 中序的位置
// 前序的位置
// |root | ------ left ------------ | ------right-------- |
// |pLeft|pLeft+1 --------         x|x+1  ----------pRight|
// 中序的位置
// | ------ left ------------| root | -------right -------|
// |iLeft--------------- pI-1| pI   | pI+1 ---------iRight|
// so x = pI - iLeft  + pLeft   //根据长度推到下即可

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	pIMap := make(map[int]int, len(inorder))
	for idx, v := range inorder {
		pIMap[v] = idx
	}

	return recuBuildTree(preorder, inorder, 0, n-1, 0, n-1, pIMap)
}

func recuBuildTree(preorder []int, inorder []int, preLeft int, preRight int, inLeft int, inRight int, piMap map[int]int) *TreeNode {
	if preLeft > preRight {
		return nil
	}

	pi := piMap[preorder[preLeft]]

	root := &TreeNode{Val: preorder[preLeft]}
	root.Left = recuBuildTree(preorder, inorder, preLeft+1, pi-inLeft+preLeft, inLeft, pi-1, piMap)
	root.Right = recuBuildTree(preorder, inorder, pi-inLeft+preLeft+1, preRight, pi+1, inRight, piMap)
	return root
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result [][]int

	que := []*TreeNode{root}
	for len(que) > 0 {
		var tmp []int
		var cnt int

		var tmpQue []*TreeNode

		for i := 0; i < len(que); i++ {
			cnt++
			item := que[i]
			tmpQue = append(tmpQue, item)
			tmp = append(tmp, item.Val)
		}

		for _, v := range tmpQue {
			if v.Left != nil {
				que = append(que, v.Left)
			}
			if v.Right != nil {
				que = append(que, v.Right)
			}
		}

		result = append(result, tmp)
		que = que[cnt:]

	}
	return result

}

func isSymmetric(root *TreeNode) bool {
	return helper(root, root)
}

func helper(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && helper(a.Left, b.Right) && helper(a.Right, b.Left)

}
