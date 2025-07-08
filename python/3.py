class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        best = 0
        buffer = []

        if len(s) == 1:
            return 1

        for char in s:
            if char not in buffer:
                buffer.append(char)
            else:
                buffer = buffer[buffer.index(char) + 1:]
                buffer.append(char)

            if len(buffer) > best:
                best = len(buffer)
        
        return best


s = Solution()
print(s.lengthOfLongestSubstring('abcc'))