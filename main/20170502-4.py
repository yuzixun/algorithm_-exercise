class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """

        m, n = len(nums1), len(nums2)

        if m > n:
            nums1, nums2, m, n = nums2, nums1, n, m

        min_num, max_num, half_num = 0, m, (n+m+1)/2

        while min_num <= max_num:
            i = (min_num + max_num)/2
            j = half_num - i

            if i < m and nums2[j-1] > nums1[i]:
                min_num = i+1
            elif i > 0 and nums1[i-1] > nums2[j]:
                max_num = i-1
            else:
                if i == 0: left = nums2[j-1]
                elif j == 0: left = nums1[i-1]
                else: left = max(nums1[i-1], nums2[j-1])

                if (m+n)%2 == 1:
                    return left
                if i == m: right = nums2[j]
                elif j == n: right = nums1[i]
                else: right = min(nums1[i], nums2[j])

                return (left + right) / 2.0

print(Solution().findMedianSortedArrays([1, 2], [3, 4, 5]))