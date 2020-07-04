class Solution(object):
    def solve(self, board):
        """
        :type board: List[List[str]]
        :rtype: void Do not return anything, modify board in-place instead.
        """

        if not any(board): return

        x = len(board[0])
        y = len(board)

        check_list = [ij for k in range(y+x) for ij in ((0, k), (k, 0), (y-1, k), (k, x-1))]
        while check_list:
            i,j = check_list.pop()
            if 0 <= i <= y-1 and 0 <= j <= x-1 and board[i][j] == 'O':
                row = list(board[i])
                row[j] = 'A'
                board[i] = ''.join(row)

                check_list += (i-1, j), (i+1, j), (i, j-1), (i, j+1)

        board[:] = [['XO'[c == 'A'] for c in row] for row in board]
        print(board)


board = ["XXXX","XOOX","XXOX","XOXX"]


print(Solution().solve(board))
