package may2020

//Definition for a binary tree node.

//TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BstFromPreorder ...
func BstFromPreorder(preorder []int) *TreeNode {
	topNode := &TreeNode{Val: preorder[0]}
	ans := topNode
	for i := 1; i < len(preorder); i++ {
		temp := topNode
		for {
			if checkNode(temp, preorder[i]) {
				if temp.Left == nil {
					temp.Left = &TreeNode{Val: preorder[i]}
					break
				} else {
					temp = temp.Left
				}
			} else {
				if temp.Right == nil {
					temp.Right = &TreeNode{Val: preorder[i]}
					break
				} else {
					temp = temp.Right
				}
			}
		}
	}
	return ans
}

func checkNode(node *TreeNode, val int) bool {
	if node.Val < val {
		return false
	}
	return true
}
