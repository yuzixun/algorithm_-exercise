class Solution(object):
    calc_hash = {
        '+': lambda a,b: a+b,
        '-': lambda a,b: a-b,
        '*': lambda a,b: a*b,
    }

    def diffWaysToCompute(self, input):
        """
        :type input: str
        :rtype: List[int]
        """

        size = len(input)

        results = []
        for x in xrange(size):
            if input[x] in ('+', '-', '*'):
                first_results = self.diffWaysToCompute(input[0:x])
                last_results = self.diffWaysToCompute(input[x+1:])
                for i in first_results:
                    for j in last_results:
                        results.append(Solution.calc_hash[input[x]](i,j))


        return results if results else [int(input)]



print(Solution().diffWaysToCompute("2*3-4*5"))
