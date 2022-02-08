package avl

import "fmt"

type node struct {
	height int
	left   *node
	right  *node
	val    Item
}

func newNodeWithVal(item Item) *node {
	return &node{
		height: 0,
		left:   nil,
		right:  nil,
		val:    item,
	}
}

func (n *node) add(item Item) *node {
	if n == nil {
		return &node{val: item, height: 1}
	}

	switch item.Less(n.val) {
	case 0:
		fmt.Println("duplicated item")
	case -1:
		n.left = n.left.add(item)
		n.updateHeight()
	case 1:
		n.right = n.right.add(item)
		n.updateHeight()
	}
	return n.balance()
}

func (n *node) delete(item Item) *node {
	// 当一个元素不存在于树，会在 nil 时 返回 nil，维持树不变，保持合理性。
	if n == nil {
		return nil
	}

	switch item.Less(n.val) {
	case 0:
		// 进行删除

		// 叶子节点直接进行删除
		if n.left == nil && n.right == nil {
			return nil
		}

		// 左右均有值时，进行替换
		// 替换为叶子节点直接进行删除
		// 替换为非叶子节点，则向下递归进行替换
		// 这里的替换只是 对 root 进行了赋值，然后递归删除 赋值的值的节点
		if n.left != nil && n.right != nil {
			if n.left.height > n.right.height {
				maxNode := n.left
				for maxNode.right != nil {
					maxNode = maxNode.right
				}
				n.val = maxNode.val
				n.left = n.left.delete(maxNode.val)

				// 删除成功时，只更新 left 的高度
				// 这个高度计算也是自下而上进行递归的
				n.left.updateHeight()
			} else {
				minNode := n.right
				for minNode.left != nil {
					minNode = minNode.left
				}
				n.val = minNode.val
				n.right = n.right.delete(minNode.val)
				n.right.updateHeight()
			}
		} else {
			if n.left != nil {
				n.val = n.left.val
				n.left = nil
			} else {
				n.val = n.right.val
				n.right = nil
			}
			n.height = 1
		}

		return n
	case -1:
		// 让左节点删除
		n.left = n.left.delete(item)
		n.updateHeight()
	case 1:
		// 让右节点删除
		n.right = n.right.delete(item)
		n.updateHeight()
	}

	// 当 case 为 1 或 -1 时，节点会收到返回的子节点
	// 子节点可能是不平衡的，需要进行平衡
	return n.balance()
}

func (n *node) balance() *node {
	// 返回该节点平衡后的结果
	var newNode *node

	if n.balanceFactor() == 2 {
		if n.left.balanceFactor() >= 0 {
			newNode = RightRotation(n)
		} else {
			newNode = LeftRightRotation(n)
		}
	} else if n.balanceFactor() == -2 {
		if n.right.balanceFactor() <= 0 {
			newNode = LeftRotation(n)
		} else {
			newNode = RightLeftRotation(n)
		}
	}

	if newNode != nil {
		newNode.updateHeight()
		return newNode
	}

	n.updateHeight()
	return n
}

func (n *node) balanceFactor() int {
	return getNodeHeight(n.left) - getNodeHeight(n.right)
}

func (n *node) updateHeight() {
	maxHeight := getNodeHeight(n.left)
	if getNodeHeight(n.right) > maxHeight {
		maxHeight = getNodeHeight(n.right)
	}
	n.height = maxHeight + 1
}

func (n *node) inOrder() {
	if n.left != nil {
		n.left.inOrder()
	}

	fmt.Printf("value: %d balance: %d\n", n.val, n.balanceFactor())
	if n.right != nil {
		n.right.inOrder()
	}
}

/*
旋转
*/

func RightRotation(n *node) *node {
	// 确定支点
	pivot := n.left

	// 保存支点的右节点
	r := pivot.right

	// 支点的右节点被 原 root 替换
	pivot.right = n

	// 原 root 的 left 为 pivot 的 right
	n.left = r

	n.updateHeight()
	pivot.updateHeight()
	return pivot
}

func LeftRotation(n *node) *node {
	// 确定支点
	pivot := n.right

	// 保留支点的左节点
	l := pivot.left

	// 将支点左节点 替换为 root
	pivot.left = n

	n.right = l

	n.updateHeight()
	pivot.updateHeight()
	return pivot
}

func LeftRightRotation(n *node) *node {
	n.left = LeftRotation(n.left)
	return RightRotation(n)
}

func RightLeftRotation(n *node) *node {
	n.right = RightRotation(n.right)
	return LeftRotation(n)
}

func getNodeHeight(n *node) int {
	if n == nil {
		return 0
	}
	return n.height
}
