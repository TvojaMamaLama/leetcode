class Solution:
    def removeDuplicates(self, nums: list[int]) -> int:
        number = -101
        new_place = 0
        for i in range(len(nums)):
            if nums[i] > number:
                number = nums[i]
                nums[new_place] = number
                new_place+=1
            else:
                nums[i] = '_'

        return new_place


s = Solution()
print(s.removeDuplicates([1,2]))
