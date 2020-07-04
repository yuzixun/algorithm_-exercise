class Solution(object):
    def updateMatrix(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: List[List[int]]
        """
        x = len(matrix[0])
        y = len(matrix)

        willCalcItems = []

        for i in range(y):
            for j in range(x):
                if matrix[i][j]:
                    matrix[i][j] = None
                else:
                    willCalcItems.append([i, j])

        while len(willCalcItems) > 0:
            i, j = willCalcItems.pop(0)

            for delta_i, delta_j in [[-1, 0], [1, 0], [0,-1], [0, 1]]:
                ci = i+delta_i
                cj = j+delta_j

                if (0 <= ci < y) and (0 <= cj < x):
                    if matrix[ci][cj] != None and matrix[ci][cj] <= matrix[i][j]+1:
                        continue
                    else:
                        matrix[ci][cj] = matrix[i][j]+1
                        willCalcItems.append([ci, cj])

        return matrix


# a = Solution().updateMatrix([[0,0,0], [0,1,0], [1,1,1]])
a = Solution().updateMatrix([[1, 0, 1, 1, 0, 0, 1, 0, 0, 1], [0, 1, 1, 0, 1, 0, 1, 0, 1, 1], [0, 0, 1, 0, 1, 0, 0, 1, 0, 0], [1, 0, 1, 0, 1, 1, 1, 1, 1, 1], [0, 1, 0, 1, 1, 0, 0, 0, 0, 1], [0, 0, 1, 0, 1, 1, 1, 0, 1, 0], [0, 1, 0, 1, 0, 1, 0, 0, 1, 1], [1, 0, 0, 0, 1, 1, 1, 1, 0, 1], [1, 1, 1, 1, 1, 1, 1, 0, 1, 0], [1, 1, 1, 1, 0, 1, 0, 0, 1, 1]])
print(a)