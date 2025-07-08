class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        buffer = {}
        for i in range(len(nums)):
            if target - nums[i] in buffer:
                return [buffer[target-nums[i]], i]
            buffer[nums[i]] = i


s = Solution()
print(s.twoSum([3,3], 6))
