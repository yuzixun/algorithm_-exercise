class Solution(object):
    def longestPalindromeSubseq(self, s):
        """
        :type s: str
        :rtype: int
        """
        d = {}

        def findIn(s):
            if s not in d:
                max_length = 0

                for i in set(s):
                    i, j = s.find(i), s.rfind(i)
                    
                    current_length = 1 if i == j else 2 + findIn(s[i+1:j])
                    max_length = max(max_length, current_length)

                d[s] = max_length
                return max_length
            else:
                return d[s]

        return findIn(s)


ss = ["bbbab", "cbbd", 'euazbipzncptldueeuechubrcourfpftcebikrxhybkymimgvldiwqvkszfycvqyvtiwfckexmowcxztkfyzqovbtmzpxojfofbvwnncajvrvdbvjhcrameamcfmcoxryjukhpljwszknhiypvyskmsujkuggpztltpgoczafmfelahqwjbhxtjmebnymdyxoeodqmvkxittxjnlltmoobsgzdfhismogqfpfhvqnxeuosjqqalvwhsidgiavcatjjgeztrjuoixxxoznklcxolgpuktirmduxdywwlbikaqkqajzbsjvdgjcnbtfksqhquiwnwflkldgdrqrnwmshdpykicozfowmumzeuznolmgjlltypyufpzjpuvucmesnnrwppheizkapovoloneaxpfinaontwtdqsdvzmqlgkdxlbeguackbdkftzbnynmcejtwudocemcfnuzbttcoew']

for s in ss:
    print(Solution().longestPalindromeSubseq(s))
