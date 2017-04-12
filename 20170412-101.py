# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def parse(self, left, right):
        if not left and not right:
            return True
        elif not left or not right:
            return False

        if left.val != right.val: return False

        return self.parse(left.left, right.right) and self.parse(left.right, right.left)

    def isSymmetric(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        if not root: return True
        return self.parse(root.left, root.right)