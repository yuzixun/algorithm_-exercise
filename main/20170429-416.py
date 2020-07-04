# -*- coding: utf-8 -*-
from collections import defaultdict

class Solution(object):
    def canPartition(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """

        nums_sum = sum(nums)

        if nums_sum % 2: return False

        half_sum = nums_sum/2

        dp = defaultdict(bool)
        dp[0] = True

        for num in nums:
            for i in reversed(xrange(num, nums_sum+1)):
                # i - num 是否可以在 某数 加上 num之后，可以达到
                dp[i] = dp[i] or dp[i-num]

        return dp[half_sum]


for arr in [[11, 1, 5,  5], [1, 2, 3, 5], [1,2,5]]:
    print(Solution().canPartition(arr))

