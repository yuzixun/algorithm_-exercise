class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class TreeGenerator(object):
    root = None
    def generate(self, unfinished, elements):
        if not unfinished:
            root = TreeNode(elements.pop(0))
            unfinished.append(root)

        # for element in elements
        while unfinished and elements:
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

        return root


# unfinished = []
# print(TreeGenerator().generate(unfinished, [1, 2,3,4,None,5,6,None,None,7]).__dict__)