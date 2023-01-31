class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        index = -1
        for i in range(len(haystack) - len(needle) + 1):
            if haystack[i] == needle[0]:
                index = i
                for k in range(len(needle)):
                    if haystack[k+i] != needle[k]:
                        index = -1
                        break
                
                if index != -1:
                    return index
        
        return index


s = Solution()
print(s.strStr(haystack = "a", needle = "a"))