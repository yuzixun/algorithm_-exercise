class Solution:
    def numOfWays(self, n: int) -> int:
        mod = int(1e9 + 7)
        s = ['ryr', 'ryg', 'rgr', 'rgy', 'yry', 'yrg', 'ygr', 'ygy', 'gry', 'grg', 'gyr', 'gyg']
        ans = {a : 1 for a in s}
        for i in range(n - 1):
            newans = {a : 0 for a in s}
            for i in s:
                for j in s:
                    if all(a != b for a, b in zip(i, j)):
                        newans[i] += ans[j]
                        newans[i] %= mod
            ans = newans
        return sum(ans.values()) % mod


Solution.new().numOfWays(5000)