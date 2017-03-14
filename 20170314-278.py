# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
# def isBadVersion(version):

class Solution(object):
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """
        earliest = 1
        lastest = n

        while earliest <= lastest:
            middle = earliest + (lastest - earliest)//2
            if isBadVersion(middle):
                lastest = middle - 1
            else:
                earliest = middle + 1

        return earliest