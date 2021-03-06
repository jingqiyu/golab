package tree

/**
 * 给定一棵二叉搜索树，请找出其中第k大的节点。



示例 1:

输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 4
示例 2:

输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 4

*/

func kthLargest(root *TreeNode, k int) int {
	var ret int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)

		if k == 0 {
			return
		}
		k--
		if k == 0 {
			ret = root.Val
		}
		dfs(root.Left)
	}
	return ret
}
