package LimitlessClassification

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	ChannelId   int    `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	Pid         int    `json:"pid"`
}

type TreeNode struct {
	Node    *Node
	Chilren []*TreeNode `json:"chilren"`
}

func (t TreeNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ChannelId   int         `json:"channel_id"`
		ChannelName string      `json:"channel_name"`
		Pid         int         `json:"pid"`
		Chilren     []*TreeNode `json:"chilren"`
	}{
		Chilren:     t.Chilren,
		ChannelId:   t.Node.ChannelId,
		ChannelName: t.Node.ChannelName,
		Pid:         t.Node.Pid,
	})
}

func listToTree(list []*Node) *TreeNode {
	nodeMap := make(map[int]*TreeNode)
	rootNode := &Node{
		1,
		"公共库",
		0,
	}
	list = append(list, rootNode)

	for _, node := range list {
		nodeMap[node.ChannelId] = &TreeNode{
			Node:    node,
			Chilren: make([]*TreeNode, 0),
		}
	}
	for _, node := range list {
		if parentTreeNode, ok := nodeMap[node.Pid]; ok {
			parentTreeNode.Chilren = append(parentTreeNode.Chilren, nodeMap[node.ChannelId])
		}
	}

	bytes, err := json.Marshal(nodeMap[1])
	fmt.Println(string(bytes), err)

	return nodeMap[1]
}

func treeToList(node *TreeNode) []*Node {
	listNode := make([]*Node, 0)
	listNode = append(listNode, node.Node)

	nodeTreeList := node.Chilren
	for len(nodeTreeList) > 0 {
		tmpNodeTreeList := make([]*TreeNode, 0) // 临时变量存储: 广度优先使用队列，深度优先使用栈
		for _, treeNode := range nodeTreeList {
			listNode = append(listNode, treeNode.Node)
			tmpNodeTreeList = append(tmpNodeTreeList, treeNode.Chilren...)
		}
		nodeTreeList = tmpNodeTreeList
	}

	bytes, err := json.Marshal(listNode)
	fmt.Println(string(bytes), err)

	return listNode
}
