class Solution(object):
    def searchMatrix(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """
        if not matrix: return False

        line_size = len(matrix[0])
        columns = 0
        rows = len(matrix) - 1

        while (columns <= line_size-1) and (rows >= 0):
            value = matrix[rows][columns]

            if value == target:
                return True
            elif value < target:
                columns += 1
            else:
                rows -= 1

        return False
