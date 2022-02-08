package avl

import (
	"testing"
)

func TestTree_Add(t *testing.T) {
	tests := []struct {
		name    string
		tr      *tree
		handler func(tr *tree)
	}{
		{
			name: "right rotation",
			tr: &tree{
				root: &node{
					val:    Int(5),
					height: 3,
					left: &node{
						val:    Int(3),
						height: 2,
						left: &node{
							val:    Int(2),
							height: 1,
						},
					},
					right: &node{
						val:    Int(6),
						height: 1,
					},
				},
			},
			handler: func(tr *tree) {
				tr.Add(Int(1))
			},
		},
		{
			name: "left rotation",
			tr: &tree{
				root: &node{
					height: 3,
					val:    Int(5),
					left: &node{
						height: 1,
						val:    Int(3),
					},
					right: &node{
						height: 2,
						val:    Int(7),
						right: &node{
							height: 1,
							val:    Int(8),
						},
					},
				},
			},
			handler: func(tr *tree) {
				tr.Add(Int(9))
			},
		},
		{
			name: "right left rotation",
			tr: &tree{
				root: &node{
					height: 3,
					val:    Int(5),
					left: &node{
						height: 1,
						val:    Int(3),
					},
					right: &node{
						height: 2,
						val:    Int(8),
						left: &node{
							height: 1,
							val:    Int(7),
						},
						right: &node{
							height: 1,
							val:    Int(9),
						},
					},
				},
			},
			handler: func(tr *tree) {
				tr.Add(Int(6))
			},
		},
		{
			name: "left right rotation",
			tr: &tree{
				root: &node{
					height: 3,
					val:    Int(10),
					left: &node{
						height: 2,
						val:    Int(5),
						left: &node{
							height: 1,
							val:    Int(3),
						},
						right: &node{
							height: 1,
							val:    Int(7),
						},
					},
					right: &node{
						height: 1,
						val:    Int(15),
					},
				},
			},
			handler: func(tr *tree) {
				tr.Add(Int(6))
			},
		},
		{
			name: "add to nil root",
			tr: &tree{
				root: nil,
			},
			handler: func(tr *tree) {
				tr.Add(Int(1))
			},
		},
		{
			name: "add duplicated item",
			tr: &tree{
				root: &node{
					height: 1,
					val:    Int(1),
				},
			},
			handler: func(tr *tree) {
				tr.Add(Int(1))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.handler(test.tr)
			test.tr.InOrder()
		})
	}
}

func TestTree_Delete(t *testing.T) {
	tests := []struct {
		name    string
		tr      *tree
		handler func(tr *tree)
	}{
		{
			name: "all nil and no rotation",
			tr: &tree{
				root: &node{
					height: 3,
					val:    Int(10),
					left: &node{
						height: 2,
						val:    Int(5),
						left: &node{
							height: 1,
							val:    Int(3),
						},
					},
					right: &node{
						height: 1,
						val:    Int(15),
					},
				},
			},
			handler: func(tr *tree) {
				tr.Delete(Int(3))
			},
		},
		{
			name: "all nil and right rotation",
			tr: &tree{
				root: &node{
					height: 3,
					val:    Int(10),
					left: &node{
						height: 2,
						val:    Int(5),
						left: &node{
							height: 1,
							val:    Int(3),
						},
					},
					right: &node{
						height: 1,
						val:    Int(15),
					},
				},
			},
			handler: func(tr *tree) {
				tr.Delete(Int(15))
			},
		},
		{
			name: "only right nil and no rotation",
			tr: &tree{
				root: &node{
					height: 3,
					val:    Int(10),
					left: &node{
						height: 2,
						val:    Int(5),
						left: &node{
							height: 1,
							val:    Int(3),
						},
					},
					right: &node{
						height: 1,
						val:    Int(15),
					},
				},
			},
			handler: func(tr *tree) {
				tr.Delete(Int(5))
			},
		},
		{
			name: "only left nil and no rotation",
			tr: &tree{
				root: &node{
					height: 3,
					val:    Int(10),
					left: &node{
						height: 1,
						val:    Int(5),
					},
					right: &node{
						height: 2,
						val:    Int(15),
						right: &node{
							height: 1,
							val:    Int(20),
						},
					},
				},
			},
			handler: func(tr *tree) {
				tr.Delete(Int(15))
			},
		},
		{
			name: "delete to nil root",
			tr: &tree{
				root: nil,
			},
			handler: func(tr *tree) {
				tr.Delete(Int(1))
			},
		},
		{
			name: "delete not exist item",
			tr: &tree{
				root: &node{
					height: 1,
					val:    Int(1),
				},
			},
			handler: func(tr *tree) {
				tr.Delete(Int(2))
			},
		},
		{
			name: "delete root",
			tr: &tree{
				root: &node{
					height: 1,
					val:    Int(1),
				},
			},
			handler: func(tr *tree) {
				tr.Delete(Int(1))
			},
		},
		{
			name: "left higher",
			tr: &tree{
				root: &node{
					height: 4,
					val:    Int(10),
					left: &node{
						height: 3,
						val:    Int(5),
						left: &node{
							height: 2,
							val:    Int(3),
							left: &node{
								height: 1,
								val:    Int(1),
							},
							right: &node{
								height: 1,
								val:    Int(4),
							},
						},
						right: &node{
							height: 1,
							val:    Int(7),
						},
					},
					right: &node{
						height: 2,
						val:    Int(15),
						right: &node{
							height: 1,
							val:    Int(20),
						},
					},
				},
			},
			handler: func(tr *tree) {
				tr.Delete(Int(5))
			},
		},
		{
			name: "right higher",
			tr: &tree{
				root: &node{
					height: 4,
					val:    Int(10),
					left: &node{
						height: 3,
						val:    Int(5),
						left: &node{
							height: 1,
							val:    Int(3),
						},
						right: &node{
							height: 2,
							val:    Int(7),
							left: &node{
								height: 1,
								val:    Int(6),
							},
							right: &node{
								height: 1,
								val:    Int(8),
							},
						},
					},
					right: &node{
						height: 2,
						val:    Int(15),
						right: &node{
							height: 1,
							val:    Int(20),
						},
					},
				},
			},
			handler: func(tr *tree) {
				tr.Delete(Int(5))
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.handler(test.tr)
			test.tr.InOrder()
		})
	}
}

func TestTree_InOrder(t *testing.T) {
	tests := []struct {
		name    string
		tr      *tree
		handler func(tr *tree)
	}{
		{
			name: "in order by nil root",
			tr: &tree{
				root: nil,
			},
			handler: func(tr *tree) {
				tr.InOrder()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.handler(test.tr)
		})
	}
}
