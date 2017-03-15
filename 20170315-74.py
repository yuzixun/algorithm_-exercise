class Solution(object):
    def searchMatrix(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """

        if not matrix:
            return False

        columns = len(matrix[0])
        rows = len(matrix)

        left = 0
        right = columns*rows - 1

        while left <= right:
            mid = (left+right)//2
            current = matrix[mid//columns][mid%columns]

            if current == target:
                return True
            elif current < target:
                left = mid + 1
            else:
                right = mid - 1

        return False

