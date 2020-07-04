class Solution(object):
    def findCircleNum(self, M):
        """
        :type M: List[List[int]]
        :rtype: int
        """
        visited = [0]*len(M)
        count = 0

        for i in range(len(M)):
            if not visited[i]:
                self.dfs(M, visited, i)
                count += 1

        return count


    def dfs(self, M, visited, i):

        for j in range(len(M)):
            if M[i][j] and not visited[j]:
                visited[j] = 1
                self.dfs(M, visited, j)






print(Solution().findCircleNum([[1,1,0], [1,1,0], [0,0,1]]))
print(Solution().findCircleNum([[1,1,0], [1,1,1], [0,1,1]]))


