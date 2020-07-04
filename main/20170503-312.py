class Solution(object):
    def maxCoins(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        nums = [1] + [num for num in nums if num > 0] + [1]
        n = len(nums)
        # n = 3
        dp = [[0] * n for _ in xrange(n)]

        for k in xrange(2,n):
            for left in xrange(0,n-k):
                right = left + k
                # print('-'*10)
                # print("left is {0}, right is {1}".format(left, right))
                for i in xrange(left+1,right):
                    dp[left][right] = max(dp[left][right], nums[left]*nums[i]*nums[right] + dp[left][i] + dp[i][right])
                    # print("nums[left]*nums[i]*nums[right] is {0}, dp[left][i] is {1}, dp[i][right] is {2}".format(nums[left]*nums[i]*nums[right], dp[left][i], dp[i][right]))
                    # print("left is {0}, right is {1}, dp is {2}".format(left, right, dp[left][right]))
                # print("result is left is {0}, right is {1}, result is {2}".format(left, right, dp[left][right]))

        return dp[0][n-1]


print(Solution().maxCoins([3, 1, 5, 8]))