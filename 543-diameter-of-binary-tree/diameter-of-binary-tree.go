/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    depth := 0

    if root != nil {
        leftDepth := 0
        rightDepth := 0

        if root.Left != nil {
            leftDepth = dfsDepth(root.Left)
        }
        if root.Right != nil {
            rightDepth = dfsDepth(root.Right)
        }

        depth = leftDepth + rightDepth
        leftChildDepth := 0
        rightChildDepth := 0
        leftChildDepth = diameterOfBinaryTree(root.Left)
        rightChildDepth = diameterOfBinaryTree(root.Right)

        if leftChildDepth > depth {
            depth = leftChildDepth
        }
        if rightChildDepth > depth {
            depth = rightChildDepth
        }
    }
    return depth
}

func dfsDepth(root *TreeNode) int {
    depth := 0
    if root != nil {
        leftDepth := dfsDepth(root.Left)
        rightDepth := dfsDepth(root.Right)

        if leftDepth > rightDepth {
            depth++
            depth += leftDepth
        } else {
            depth++
            depth += rightDepth
        }
    }
    return depth
}