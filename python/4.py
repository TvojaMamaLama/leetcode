class Solution:
    def findMedianSortedArrays(self, nums1: list[int], nums2: list[int]) -> float:
        first_len = len(nums1)
        second_len = len(nums2)
        point1 = 0
        point2 = 0
        result_mas = []

        while True:
            if point1 >= first_len:
                result_mas += nums2[point2:]
                break
            
            if point2 >= second_len:
                result_mas += nums1[point1:]
                break

            if nums1[point1] >= nums2[point2]:
                result_mas.append(nums2[point2])
                point2 += 1
            else:
                result_mas.append(nums1[point1])
                point1 += 1
        
        res_len = len(result_mas)
        if res_len % 2 != 0:
            return result_mas[int(res_len / 2)]
        else:
            return (result_mas[int(res_len / 2 - 1)] + result_mas[int(res_len / 2)]) / 2



s = Solution()
print(s.findMedianSortedArrays([1,2, 90, 91], [2, 3, 4, 100]))
