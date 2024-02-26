/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    pList := traverse(p)
    qList := traverse(q)

    if len(pList) != len(qList) {
        return false
    }

    for i := 0; i < len(pList); i++ {
        pNode := pList[i]
        qNode := qList[i]

        if pNode != nil && qNode != nil {
            if pNode.Val != qNode.Val {
                return false
            }
        } else if pNode == nil {
            if qNode != nil {
                return false
            }
        } else if qNode == nil {
            return false
        }
    }

    return true
}

func traverse(a *TreeNode) []*TreeNode {
    nodeList := []*TreeNode{}
    nodeQueue := []*TreeNode{}

    nodeQueue = append(nodeQueue, a)
    for len(nodeQueue) > 0 {
        currentNode := nodeQueue[0]
        nodeQueue = nodeQueue[1:]
        nodeList = append(nodeList, currentNode)
        if currentNode != nil {
            nodeQueue = append(nodeQueue, currentNode.Left)
            nodeQueue = append(nodeQueue, currentNode.Right)
        }
    }

    return nodeList
}