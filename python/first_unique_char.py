class Solution:
    def firstUniqChar(self, s: str) -> int:
        counter = {}
        for char in s:
            if char in counter:
                counter[char] += 1
            else:
                counter[char] = 1
        
        for idx, char in enumerate(s):
            if counter[char] == 1:
                return idx

        return -1
            


s = Solution()
print(s.firstUniqChar('leetcode'))