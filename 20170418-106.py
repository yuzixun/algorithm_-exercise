# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution(object):
    def buildTree(self, inorder, postorder):
        """
        :type inorder: List[int]
        :type postorder: List[int]
        :rtype: TreeNode
        """
        if not inorder or not postorder:
            return None

        root_val = postorder.pop()

        root = TreeNode(root_val)
        root_index_of_inorder = inorder.index(root_val)

        root.right = self.buildTree(inorder[root_index_of_inorder+1:], postorder)
        root.left = self.buildTree(inorder[:root_index_of_inorder], postorder)

        return root


# Solution().buildTree([-1], [-1])
Solution().buildTree([2,1,3], [2,3,1])
print(Solution().buildTree([2,1], [2,1]).__dict__)