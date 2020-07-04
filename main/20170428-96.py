class Solution(object):
    def numTrees(self, n):
        """
        :type n: int
        :rtype: int
        """
        dp = {}
        dp[0] = 1
        dp[1] = 1
        for x in xrange(2, n+1):
            dp[x] = 0
            for i in xrange(1, x+1):
                dp[x] += dp[i-1] * dp[x-i]
            # print("{0}, {1}".format(x, dp[x]))

        return dp[n]


for n in [3]:
    print(Solution().numTrees(n))