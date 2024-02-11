/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // Loop through each node while next != nil
    // Add both value, store the addition for the next iteration

    resultNode := &ListNode{}
    currentNode := resultNode
    addition := 0

    for l1 != nil && l2 != nil {
        currentNode.Val = (l1.Val + l2.Val + addition) % 10
        addition = (l1.Val + l2.Val + addition) / 10

        l1 = l1.Next
        l2 = l2.Next
        if l1 != nil && l2 != nil {
            currentNode.Next = &ListNode{}
            currentNode = currentNode.Next
        }
    }

    var additionalNode *ListNode
    if l1 != nil {
        additionalNode = l1
    } else if l2 != nil {
        additionalNode = l2
    }

    for additionalNode != nil {
        currentNode.Next = &ListNode{}
        currentNode = currentNode.Next
        currentNode.Val = (additionalNode.Val + addition) % 10
        addition = (additionalNode.Val + addition) / 10
        additionalNode = additionalNode.Next
    }

    if addition > 0 {
        currentNode.Next = &ListNode{
            Val: addition,
        }
    }

    return resultNode
}