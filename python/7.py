class Solution:
    def reverse(self, x: int) -> int:
        new = str(x)[::-1].strip('-').strip('0')
        new = new or 0
        if int(new) > 2147483648:
            return 0

        if x < 0:
            new = '-' + new

        return int(new)


s = Solution()
print(s.reverse(0))