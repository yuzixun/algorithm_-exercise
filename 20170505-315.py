class Solution(object):

    def countSmaller(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        result = [0]*len(nums)
        def sort_and_count_reverse(elements):
            half = len(elements)/2

            if half:
                left, right = sort_and_count_reverse(elements[:half]), sort_and_count_reverse(elements[half:])

                for i in range(len(elements))[::-1]:
                    if not right or left and left[-1][1] > right[-1][1]:
                        # right 是经过从小到大排序的，
                        # 所以 遇到current_right < current_left的情况
                        # 可以确定，right中剩下的数，都比current_left小
                        result[left[-1][0]] += len(right)
                        elements[i] = left.pop() # 归并排序
                    else:
                        elements[i] = right.pop() # 归并排序

            return elements

        sort_and_count_reverse(list(enumerate(nums)))
        return result


print(Solution().countSmaller([5,2,6,1]))
