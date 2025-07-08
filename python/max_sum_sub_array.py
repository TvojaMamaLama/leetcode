def max_sub_array(nums: list[int]) -> int:
    max_sum = 0
    buffer = 0

    for num in nums:
        buffer = max(buffer + num, num)

        if buffer > max_sum:
            max_sum = buffer

    return max_sum


print(max_sub_array([-2, 1, -3, 4, -1, 2, 1, -5, 4]))
