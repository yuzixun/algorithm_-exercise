# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def findBottomLeftValue(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        if not root: return 0

        queue = [root]

        while queue:
            root = queue.pop(0)
            queue += filter(None, (root.right, root.left))

        return root.val



def generate(unfinished, elements):
    # for element in elements
    while len(unfinished) and elements:
        current_node = unfinished.pop(0)

        # print('------')
        if elements:
            left_node = TreeNode(elements.pop(0))

            current_node.left = left_node
            if left_node.val: unfinished.append(left_node)

        if elements:
            right_node = TreeNode(elements.pop(0))

            current_node.right = right_node
            if right_node.val: unfinished.append(right_node)

        # print '------'
        # print(current_node.__dict__)
        # if current_node.left: print(current_node.left.__dict__)
        # if current_node.right: print(current_node.right.__dict__)


root = TreeNode(1)
unfinished = [root]
generate(unfinished, [2,3,4,None,5,6,None,None,7])
# print(root.__dict__)
# tree =
# elements = [1,2,3,4,None,5,6,None,None,7]
# for element in elements


print(Solution().findBottomLeftValue(root))