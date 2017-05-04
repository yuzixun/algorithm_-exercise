class Solution(object):
    def reverseStr(self, s, k):
        """
        :type s: str
        :type k: int
        :rtype: str
        """

        s_len = len(s)

        result = ''
        l, r = 0, 2*k
        while True:
            result += s[l:l+k][::-1]
            result += s[l+k:r]

            if r > s_len:
                break

            l += 2*k
            r += 2*k

        return result


s = "abcdefg"
k = 2
# bacdfeg
print(Solution().reverseStr(s, k))