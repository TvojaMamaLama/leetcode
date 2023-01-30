class Solution:
    def isPalindrome(self, s: str) -> bool:
        clean_str = [char.lower() for char in s if char.isalpha()]
        if clean_str == clean_str[::-1]:
            return True

        return False
        

s = Solution()
print(s.isPalindrome('A man, a plan, a canal: Panama'))