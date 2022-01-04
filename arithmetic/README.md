## ARITHMETIC

### Binary Tree

包含二叉树相关的算法题

#### PreorderTraversal

前序遍历二叉树

#### PreorderTraversalII 

前序遍历二叉树 非递归的方式

#### InorderTraversal

中序遍历二叉树 非递归

#### PostorderTraversal

后序遍历 非递归

#### PreorderTraversalIII

DFS 深度优先遍历 分治法 递归和回溯

#### LevelOrder

BFS 广度优先遍历

#### LevelOrderII

BFS 广度优先遍历 改进版

#### LevelOrderBottom

自下而上BFS, 使用LeverOrder的方法，把结果反转即可

#### ZigzagLevelOrder

自上而下, Z字形遍历, 增加一个toggle变量，控制level是否需要反转即可

#### MaxDepth

二叉树的最大深度  递归 分治法

#### IsBalanced

判断平衡二叉树 分治法 返回值的二义性

#### MaxPathSum

二叉树中的最大路径和，路径每到一个节点，有 3 种选择：
1. 停在当前节点。
2. 走到左子节点。
3. 走到右子节点。  

使用dfs, 每次更新maxSum，然后每次返回该节点的最大路径，并且负数时返回0

#### LowestCommonAncestor

查找最近公共祖先， 通过递归对二叉树进行后序遍历，当遇到节点 p 或 q 时返回。  
从底至顶回溯，当节点 p, q 在节点 root 的异侧时，节点 root 即为最近公共祖先，则向上返回 root。

#### IsValidBST

判断二叉搜索树， 左子树返回一个最大的数, 右子树返回一个最小的数, 然后和根节点比较

#### IsValidBSTII 

验证二叉搜索树，使用中序遍历二叉树，然后判断排序是否正确

#### InsertIntoBST 

插入二叉搜索树，将新插入的节点插入到叶子节点即可

### Sort

#### MergeSort

归并排序 分治法

#### QuickSort

快速排序 分治法
关键是如何进行分区

### Links

#### DeleteDuplicates

给定一个按升序排序的链表  
删除链表中重复的元素  
错误记录:
1. 没有加上for循环
2. 没有加入当前节点p向后移动的逻辑，导致无限循环
3. 需要等重复节点全部删除完，再移到下一个节点

### math

#### Sqrt

使用牛顿法不断推算平方根