import collections

class Solution(object):
    def findItinerary(self, tickets):
        """
        :type tickets: List[List[str]]
        :rtype: List[str]
        """
        result = ['JFK']
        grouped_tickets = collections.defaultdict(list)

        for source, destination in sorted(tickets):
            grouped_tickets[source].append(destination)

        def visit(airport):
            while grouped_tickets[airport]:
                current = grouped_tickets[airport].pop(0)
                visit(current)
                result.insert(1, current)

        visit(result[0])
        return result



# tickets = [["MUC", "LHR"], ["JFK", "MUC"], ["SFO", "SJC"], ["LHR", "SFO"]]
tickets = [["JFK","KUL"],["JFK","NRT"],["NRT","JFK"]]
# tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
print(Solution().findItinerary(tickets))

# {'JFK': ['MUC'], 'SFO': ['SJC'], 'LHR': ['SFO'], 'MUC': ['LHR']})