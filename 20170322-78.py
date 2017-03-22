class Solution(object):
    def subsets(self, nums):
        res = [[]]

        for num in sorted(nums):
            res += [item+[num] for item in res]
            print(res)

        return res


[item+[2] for item in [[], [1]]]


print(Solution().subsets([1, 2, 3]))