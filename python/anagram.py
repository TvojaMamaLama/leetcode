class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if sorted(s) == sorted(t):
            return True
        return False


s = Solution()
print(s.isAnagram('aba', 'baa'))
