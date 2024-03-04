def max_product(nums: list):
    if len(nums) == 2:
        return nums[0] * nums[1]

    pos1, pos2, neg1, neg2 = 0, 0, 0, 0
    for i in range(len(nums)):
        if nums[i] > 0:
            if nums[i] >= pos1:
                pos1, pos2 = nums[i], pos1
            elif nums[i] > pos2:
                pos2 = nums[i]

        if nums[i] < 0:
            if nums[i] <= neg1:
                neg1, neg2 = nums[i], neg1
            elif nums[i] < neg2:
                neg2 = nums[i]

    return max(pos1 * pos2, neg1 * neg2)


print(max_product([-2, 1, -3, 4, -1, 2, 1, -5, 4]))
