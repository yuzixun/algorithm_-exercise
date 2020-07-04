from tree_generator import TreeGenerator
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def check(self, root, min, max):
        if not root: return True

        if (min and root.val <= min.val) or (max and root.val >= max.val):
            return False
        else:
            return self.check(root.left, min, root) and self.check(root.right, root, max)

    def isValidBST(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        return self.check(root, None, None)






# root = TreeGenerator().generate([], [2,1,3,1,3,2,4])
# root = TreeGenerator().generate([], [1,2,3])
# root = TreeGenerator().generate([], [1,1])
root = TreeGenerator().generate([], [10,5,15,None,None,6,20])
print(Solution().isValidBST(root))