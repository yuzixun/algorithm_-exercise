from tree_generator import TreeGenerator

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    prev = None

    def __init__(self):
        Solution.prev = None


    def flatten(self, root):
        """
        :type root: TreeNode
        :rtype: void Do not return anything, modify root in-place instead.
        """

        if not root: return None

        self.flatten(root.right)
        self.flatten(root.left)

        root.right = Solution.prev
        root.left = None

        Solution.prev = root
        print(root.__dict__)



root = TreeGenerator().generate([], [1,2])
print(Solution().flatten(root))
print(root.__dict__)
print(root.right.__dict__)
