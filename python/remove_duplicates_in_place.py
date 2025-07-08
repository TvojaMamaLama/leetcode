def remove_duplicates(nums: list) -> int:
    if not nums:
        return 0

    pointer = 0
    for i in range(1, len(nums)):
        if nums[pointer] < nums[i]:
            nums[pointer+1] = nums[i]
            pointer += 1

    return pointer + 1



nums = [1,1,2,2,2,3,3,3,3,3]
print(remove_duplicates(nums))
print(nums)

