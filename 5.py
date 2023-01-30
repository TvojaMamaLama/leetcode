class Solution:
    def longestPalindrome(self, s: str) -> str:
        current_longer = (0, 1)

        for i in range(1, len(s)):
            letter_polindrom = self.check_palindrome(s, i-1, i+1)
            middle_polindrome = self.check_palindrome(s, i-1, i)

            longer = max(middle_polindrome, letter_polindrom, key= lambda x: x[1] - x[0])
            current_longer = max(longer, current_longer, key= lambda x: x[1] - x[0])
        
        return s[current_longer[0]:current_longer[1]]
        
    def check_palindrome(self, s: str, left: int, right: int):
        while left >= 0 and right < len(s) and s[left] == s[right]:
            left -= 1
            right += 1
        
        return (left + 1, right)


s = Solution()
print(s.longestPalindrome('abacddcc'))