# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def isPalindrome(self, head: ListNode) -> bool:
        slow, fast, new_head = head, head, None
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next

        fast = head
        new_head = slow
        slow = slow.next
        new_head.next = None

        while slow:
            temp = slow.next
            slow.next = new_head
            new_head = slow
            slow = temp

        while fast and new_head:
            if fast.val != new_head.val:
                return False

            fast = fast.next
            new_head = new_head.next

        return True


list1 = ListNode(1, ListNode(3, ListNode(5, ListNode(3, ListNode(1)))))
print(Solution().isPalindrome(list1))

# while True:
#     print(node.val)
#     node = node.next
#
#     if node is None:
#         break
