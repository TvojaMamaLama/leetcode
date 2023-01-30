class Solution:
    def singleNumber(self, nums: list[int]) -> int:
        counter = 0
        for i in nums:
            counter ^= i

        return counter


s = Solution()
print(s.singleNumber([1,1,2,2,3]))