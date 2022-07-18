/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode mergeTwoLists(ListNode list1, ListNode list2) {
        ListNode root = null;
        if (list1 == null && list2 == null) return root;
        
        if (list1 != null) {
            if (list2 == null) {
                root = list1;
                root.next = this.mergeTwoLists(list1.next, list2);
            } else if (list1.val <= list2.val) {
                root = list1;
                root.next = this.mergeTwoLists(list1.next, list2);
            }
        }
        if (list2 != null) {
            if (list1 == null) {
                root = list2;
                root.next = this.mergeTwoLists(list1, list2.next);
            } else if (list2.val < list1.val) {
                root = list2;
                root.next = this.mergeTwoLists(list1, list2.next);
            }
        }
        
        return root;
    }
}