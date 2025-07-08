# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def hasCycle(self, head: ListNode) -> bool:
        if head is None:
            return False

        fast = slow = head
        while fast and fast.next is not None:
            fast = fast.next.next
            slow = slow.next

            if fast == slow:
                return True

        return False


node = ListNode(4)
list1 = ListNode(1, ListNode(3, node))
node.next = list1

print(Solution().hasCycle(list1))

# while True:
#     print(node.val)
#     node = node.next
#
#     if node is None:
#         break
