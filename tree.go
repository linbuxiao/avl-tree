package avl

// 我猜 avl 与 bst 没什么不同，只是要在插入和删除后进行一次自平衡。

// 所以先使用 bst 进行实现

// 很大的启发时：之前我认为对一个节点操作需要找到它的父节点后，通过 parent-node + node 进行操作，但其实使用递归可以直接进行返回赋值。
// 永远站在 父节点上可以避免纠结于子节点

type tree struct {
	root *node
}

func (tr *tree) Add(item Item) *tree {
	if tr.root == nil {
		tr.root = newNodeWithVal(item)
	} else {
		tr.root = tr.root.add(item)
	}
	return tr
}

func (tr *tree) Delete(item Item) *tree {
	// 删除是一个递归移动被删除节点的过程，直到它转移至叶子节点
	// 当具备叶子节点时，可以直接替换 在不具备叶子节点条件时，只可慢慢移动
	if tr.root == nil {
		return nil
	}
	if item.Less(tr.root.val) == 0 {
		tr.root = nil
		return tr
	}

	tr.root = tr.root.delete(item)
	return tr
}

func (tr *tree) InOrder() {
	if tr.root == nil {
		return
	}

	tr.root.inOrder()
}
