class Solution(object):

    def findMin(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """

        left = 0
        right = len(nums) - 1

        while left < right:
            if nums[left] < nums[right]:
                return nums[left]

            mid = (left + right) // 2
            if nums[mid] >= nums[left]:
                left = mid + 1
            else:
                right = mid

        return nums[left]


print(Solution().findMin([1,2,3,4,5,6,7,0,]))
print(Solution().findMin([7,0,1,2,3,4,5,6,]))