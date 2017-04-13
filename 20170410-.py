# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def getResult(self, root, result, level):
        if not root: return None
        if len(result) <= level: result.insert(0, [])

        result[-(level+1)].append(root.val)

        if root.left: self.getResult(root.left, result, level+1)
        if root.right: self.getResult(root.right, result, level+1)



    def levelOrderBottom(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        result = []
        self.getResult(root, result, 0)
        return result
