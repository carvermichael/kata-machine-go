package misc

var Tree = &BinaryTreeNode{
	val: 20,
	left: &BinaryTreeNode{
		val: 10,
		left: &BinaryTreeNode{
			val:  5,
			left: nil,
			right: &BinaryTreeNode{
				val:   7,
				left:  nil,
				right: nil,
			},
		},
		right: &BinaryTreeNode{
			val:   15,
			left:  nil,
			right: nil,
		},
	},
	right: &BinaryTreeNode{
		val: 50,
		left: &BinaryTreeNode{
			val: 30,
			left: &BinaryTreeNode{
				val:   29,
				left:  nil,
				right: nil,
			},
			right: &BinaryTreeNode{
				val:   45,
				left:  nil,
				right: nil,
			},
		},
		right: &BinaryTreeNode{
			val:   100,
			left:  nil,
			right: nil,
		},
	},
}

var Tree2 = &BinaryTreeNode{
	val: 20,
	right: &BinaryTreeNode{
		val:   50,
		right: nil,
		left: &BinaryTreeNode{
			val: 30,
			right: &BinaryTreeNode{
				val: 45,
				right: &BinaryTreeNode{
					val:   49,
					left:  nil,
					right: nil,
				},
				left: nil,
			},
			left: &BinaryTreeNode{
				val:   29,
				right: nil,
				left: &BinaryTreeNode{
					val:   21,
					right: nil,
					left:  nil,
				},
			},
		},
	},
	left: &BinaryTreeNode{
		val: 10,
		right: &BinaryTreeNode{
			val:   15,
			right: nil,
			left:  nil,
		},
		left: &BinaryTreeNode{
			val: 5,
			right: &BinaryTreeNode{
				val:   7,
				right: nil,
				left:  nil,
			},
			left: nil,
		},
	},
}
