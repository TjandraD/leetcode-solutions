/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    nodeOutput := dfsBottomLeft(root, 1)
    return nodeOutput.Val
}

type NodeOutput struct {
    Val int
    Depth int
}

func dfsBottomLeft(root *TreeNode, depth int) NodeOutput {
    leftOutput := NodeOutput{
        Val: root.Val,
        Depth: depth,
    }
    rightOutput := NodeOutput{
        Val: root.Val,
        Depth: depth,
    }

    if root.Left != nil {
        leftOutput = dfsBottomLeft(root.Left, depth+1)
    }
    if root.Right != nil {
        rightOutput = dfsBottomLeft(root.Right, depth+1)
    }

    if leftOutput.Depth >= rightOutput.Depth {
        return leftOutput
    }
    return rightOutput
}