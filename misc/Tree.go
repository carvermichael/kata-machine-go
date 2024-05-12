package misc

var Tree = &BinaryTreeNode[int]{
	Val: 20,
	Left: &BinaryTreeNode[int]{
		Val: 10,
		Left: &BinaryTreeNode[int]{
			Val:  5,
			Left: nil,
			Right: &BinaryTreeNode[int]{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BinaryTreeNode[int]{
			Val:   15,
			Left:  nil,
			Right: nil,
		},
	},
	Right: &BinaryTreeNode[int]{
		Val: 50,
		Left: &BinaryTreeNode[int]{
			Val: 30,
			Left: &BinaryTreeNode[int]{
				Val:   29,
				Left:  nil,
				Right: nil,
			},
			Right: &BinaryTreeNode[int]{
				Val:   45,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &BinaryTreeNode[int]{
			Val:   100,
			Left:  nil,
			Right: nil,
		},
	},
}

var Tree2 = &BinaryTreeNode[int]{
	Val: 20,
	Right: &BinaryTreeNode[int]{
		Val:   50,
		Right: nil,
		Left: &BinaryTreeNode[int]{
			Val: 30,
			Right: &BinaryTreeNode[int]{
				Val: 45,
				Right: &BinaryTreeNode[int]{
					Val:   49,
					Left:  nil,
					Right: nil,
				},
				Left: nil,
			},
			Left: &BinaryTreeNode[int]{
				Val:   29,
				Right: nil,
				Left: &BinaryTreeNode[int]{
					Val:   21,
					Right: nil,
					Left:  nil,
				},
			},
		},
	},
	Left: &BinaryTreeNode[int]{
		Val: 10,
		Right: &BinaryTreeNode[int]{
			Val:   15,
			Right: nil,
			Left:  nil,
		},
		Left: &BinaryTreeNode[int]{
			Val: 5,
			Right: &BinaryTreeNode[int]{
				Val:   7,
				Right: nil,
				Left:  nil,
			},
			Left: nil,
		},
	},
}
