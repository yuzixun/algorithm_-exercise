# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def largestValues(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        nodes_list = [root]
        result = []
        while any(nodes_list):
            result.append(max(node.val for node in nodes_list))
            nodes_list = [kid for node in nodes_list for kid in (node.right, node.left) if kid]
        return result

def generate(unfinished, elements):
    # for element in elements
    while unfinished and elements:
        current_node = unfinished.pop(0)

        if elements:
            left_node = TreeNode(elements.pop(0))

            current_node.left = left_node
            if left_node.val: unfinished.append(left_node)

        if elements:
            right_node = TreeNode(elements.pop(0))

            current_node.right = right_node
            if right_node.val: unfinished.append(right_node)

        print('-------')
        print(current_node.__dict__)
        if current_node.left:  print(current_node.left.__dict__)
        if current_node.right: print(current_node.right.__dict__)


root = TreeNode(1)
unfinished = [root]
generate(unfinished, [3,2,5,3,None,9])

Solution().largestValues(root)