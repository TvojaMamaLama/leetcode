package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	k, j := m-1, n-1
	for i := len(nums1) - 1; i >= 0; i-- {
		if k < 0 {
			nums1[i] = nums2[j]
			j--
			continue
		}

		if j < 0 {
			nums1[i] = nums1[k]
			k--
			continue
		}

		if nums1[k] >= nums2[j] {
			nums1[i] = nums1[k]
			k--
		} else {
			nums1[i] = nums2[j]
			j--
		}
	}
}