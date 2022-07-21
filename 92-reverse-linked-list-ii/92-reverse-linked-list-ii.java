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
    public ListNode reverseBetween(ListNode head, int left, int right) {
        if (left == right) return head;
        
        ListNode headPointer = new ListNode(0);
        headPointer.next = head;
        ListNode preSwap = headPointer;
        
        for (int i = 0; i < left - 1; i++) preSwap = preSwap.next;
        
        ListNode toSwap = preSwap.next;
        ListNode nextSwap = toSwap.next;
        
        for (int i = 0; i < right - left; i++) {
            toSwap.next = nextSwap.next;
            nextSwap.next = preSwap.next;
            preSwap.next = nextSwap;
            nextSwap = toSwap.next;
        }
        
        return headPointer.next;
    }
}