class Solution:
    def containsDuplicate(self, nums: list[int]) -> bool:
        s = set(nums)
        for i in nums:
            if i in s:
                return True
            s.add(i)

        return False


s = Solution()
print(s.containsDuplicate([1,2,3,4]))