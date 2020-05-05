package Tree

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// 程序: https://blog.csdn.net/qq_41682302/article/details/94636008
// 动画: http://btv.melezinek.cz/binary-search-tree.html
// 前序遍历(第一次遇到该结点时输出):
//  1. 输出当前节点
//  2. 遍历左子树，一直向左下走，直到走到最左下的叶子结点
//  3. 遍历完左子树后，从栈中弹出一个节点，回撤到当前叶子的双亲节点（借用栈结构，先进后出实现）
//  4. 遍历右子树（指定当前节点为栈回撤出来节点的右节点，重复1~3步操作）
func preOrder(node *Node) (array []int) {
	if node == nil {
		return array
	}
	rootNode := node
	stack := make([]*Node, 0)
	// 回撤栈不为空或者当前节点不是空子树时说明未遍历完成
	for len(stack) != 0 || rootNode != nil {
		// 前序遍历对针对每个结点，都尽量往左去寻找，直到走到最左下的叶子结点
		// 每遍历一个节点时，我们将节点打印出来，并且将当前节点入栈(以供后续变量右子树时回撤)
		for rootNode != nil {
			// 先访问
			array = append(array, rootNode.Value)
			fmt.Println(array)
			// 再入栈
			stack = append(stack, rootNode)
			rootNode = rootNode.Left
		}

		// 上面的for执行完了说明左子树已经遍历到了最左子节点
		if len(stack) != 0 {
			// 回撤，指定遍历右子树
			rootNode = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}

	return array
}

// 前序遍历(第二次遇到该结点时输出):
//  1. 遍历左子树，一直向左下走，直到走到最左下的叶子结点
//  2. 遍历完左子树后，从栈中弹出一个节点，回撤到当前叶子的双亲节点（借用栈结构，先进后出实现）
//  3. 输出节点
//  4. 遍历右子树（指定当前节点为栈回撤出来节点的右节点，重复1~3步操作）
func inOrder(node *Node) (array []int) {
	if node == nil {
		return array
	}
	rootNode := node
	stack := make([]*Node, 0)
	// 回撤栈不为空或者当前节点不是空子树时说明未遍历完成
	for len(stack) != 0 || rootNode != nil {
		// 中遍历对每一棵树而言，都一直向左下走，直到走到最左下的叶子结点
		for rootNode != nil {
			stack = append(stack, rootNode)
			rootNode = rootNode.Left
		}

		// 每遍历一个节点时，我们将节点打印出来，并且将当前节点入栈(以供后续变量右子树时回撤)
		// 上面的for执行完了说明左子树已经遍历到了最左子节点
		if len(stack) != 0 {
			// 回撤，指定遍历右子树
			pop := stack[len(stack)-1]
			array = append(array, pop.Value)
			rootNode = pop.Right
			stack = stack[:len(stack)-1]
		}
	}

	return array
}

func postOrder(node *Node) (array []int) {
	if node == nil {
		return array
	}

	stack := make([]*Node, 0)
	stack = append(stack, node)

	for len(stack) != 0 {
		root := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		if root.Right != nil {
			stack = append(stack, root.Right)
		}

		// 按照后续遍历的规则，node节点一定是最后访问的
		// 我们可以使用逆向队列的思维，先反着放，最后在倒过来
		fmt.Println(root.Value, root)
		array = append(array, root.Value)
	}

	// 翻转结果队列
	tmpArray := make([]int, 0)
	for i := len(array) - 1; i >= 0; i-- {
		tmpArray = append(tmpArray, array[i])
	}

	return tmpArray
}
