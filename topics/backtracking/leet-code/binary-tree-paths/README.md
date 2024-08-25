# 257. Binary Tree Paths
- Easy, Topics, Companies

Given the root of a binary tree, return all root-to-leaf paths in any order.

A leaf is a node with no children.

### Example 1:

```
Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]
```

### Example 2:

```
Input: root = [1]
Output: ["1"]
```

### Constraints:

- The number of nodes in the tree is in the range [1, 100].
- -100 <= Node.val <= 100


## Solution

- Enumerate all DFS paths
- For what I know so far the DFS will use backtracking to traverse the hole tree

1. Choice: either go left or right at any given node
2. Constraint: If there is no other node to visit from the current node, then go back
3. Goal: If its a valida path store that node into the path
