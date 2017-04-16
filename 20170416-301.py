class Solution(object):
    def removeInvalidParentheses(self, s):
        """
        :type s: str
        :rtype: List[str]
        """
        check_list = [s]
        result = []

        while check_list:
            result = filter(self.is_valid, check_list)

            if result:
                return result
            else:
                check_list = set([string[0:i]+string[i+1:] for string in check_list for i in range(len(string))])



    def is_valid(self, string):
        counter = 0
        for s in string:
            if s == '(':
                counter += 1
            elif s == ')':
                counter -= 1
                if counter < 0: return False
            else:
                continue

        return counter == 0

# print(Solution().removeInvalidParentheses("()())()"))
print(Solution().removeInvalidParentheses("n"))