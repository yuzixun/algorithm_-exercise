class Solution(object):
    def countBits(self, num):
        """
        :type num: int
        :rtype: List[int]
        """
        results = [0]
        offset = 1
        # for x in xrange(1, num+1): # 229 ms
        for x in range(1, num+1): # 185 ms
            if offset*2 == x: offset *= 2

            results.append(results[x-offset] + 1)

        return results

print(Solution().countBits(5))