class Solution(object):
    def addOperators(self, num, target):
        """
        :type num: str
        :type target: int
        :rtype: List[str]
        """

        def generate(target, pos, sym, value):

            results = []

            for i in xrange(pos, len(num)):
                if i > pos and num[pos] == '0':
                    break

                if i == len(num) - 1:
                    if sym * value * int(num[pos:i+1]) == target:
                        results.extend([num[pos:i+1]])
                    break

                current_value = value * sym * (int(num[pos:i+1]))

                results.extend([num[pos:i+1] + '+' + e for e in generate(target - current_value, i+1, 1, 1)])
                results.extend([num[pos:i+1] + '-' + e for e in generate(target - current_value, i+1, -1, 1)])
                results.extend([num[pos:i+1] + '*' + e for e in generate(target, i+1, 1, current_value)])

            return results

        return generate(target, 0, 1, 1)

print(Solution().addOperators("123", 6))
