class Solution:
    def intersect(self, nums1: list[int], nums2: list[int]) -> list[int]:
        first = 0
        second = 0
        answer = []
        nums2.sort()
        nums1.sort()

        while first < len(nums1) and second < len(nums2):
            if nums1[first] == nums2[second]:
                answer.append(nums1[first])
                first += 1
                second += 1
            elif nums1[first] > nums2[second]:
                second += 1
            else:
                first += 1
        
        return answer


s = Solution()
print(s.intersect([1,2,2,1], [4,2,2,4]))
