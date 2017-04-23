from collections import defaultdict
# class Solution(object):
    # def findTargetSumWays(self, nums, S):
    #     """
    #     :type nums: List[int]
    #     :type S: int
    #     :rtype: int
    #     """

    #     nums_sum = sum(nums)

    #     if (S+nums_sum)%2 == 0:
    #         result = self.sub_sum(nums, int((S+nums_sum)/2))
    #     else:
    #         result = 0

    #     return nums_sum < S or result

    # def sub_sum(self, nums, s):
    #     dp = {}
    #     dp = defaultdict(lambda:0, dp)
    #     dp[0] = 1

    #     for n in nums:
    #         for i in range(s, n-1, -1):
    #             dp[i] += dp[i-n]
    #             print("n is {0}, i is {1}, dp is {2}, s is {3}".format(n, i, dp, s))

    #         # print('---------')
    #         # print(dp)

    #     return dp[s]
class Solution(object):
    def findTargetSumWays(self, nums, S):
        from collections import defaultdict
        memo = {0: 1}
        for x in nums:
            # print('-----')
            # print(x)
            m = defaultdict(int)
            for s, n in memo.iteritems():
                m[s + x] += n
                m[s - x] += n
            memo = m
            print(memo)
        return memo[S]


# nums = [1, 1, 3, 1, ]
# S = 3

nums = [1,2,7,9,981]
S = 1000000000

print(Solution().findTargetSumWays(nums, S))