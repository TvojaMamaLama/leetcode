class Solution:
    def merge(self, nums1: list[int], m: int, nums2: list[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        pointer1, pointer2, pointer3 = m-1, n-1, m+n-1
        while pointer2 >= 0:
            if pointer1 >= 0 and nums1[pointer1] > nums2[pointer2]:
                nums1[pointer3] = nums1[pointer1]
                pointer1 -= 1
            else:
                nums1[pointer3] = nums2[pointer2]
                pointer2 -= 1

            pointer3 -= 1

        print(nums1)


s = Solution().merge([1,2,3,0,0,0], 3, [2,5,6], 3)
