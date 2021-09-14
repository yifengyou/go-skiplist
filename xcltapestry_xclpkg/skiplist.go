package algorithm

//SkipList
//author:Xiong Chuan Liang
//date:2014-1-28
//"github.com/xcltapestry/xclpkg/algorithm"

import (
	//"fmt"
	"fmt"
	"math/rand"
)

const SKIPLIST_MAXLEVEL = 32 //8
const SKIPLIST_P = 4

type Node struct {
	Forward []Node
	Value   interface{}
}

type SkipList struct {
	Header *Node
	Level  int
}

func NewNode(v interface{}, level int) *Node {
	// Value填充为level
	return &Node{Value: v, Forward: make([]Node, level)}
}

func NewSkipList() *SkipList {
	return &SkipList{Level: 1, Header: NewNode(0, SKIPLIST_MAXLEVEL)}
}

func (skipList *SkipList) Insert(key int) {
	update := make(map[int]*Node)
	node := skipList.Header

	for i := skipList.Level - 1; i >= 0; i-- {
		for {
			if node.Forward[i].Value != nil && node.Forward[i].Value.(int) < key {
				node = &node.Forward[i]
			} else {
				break
			}
		}
		// 前序节点位置
		update[i] = node
	}
	level := skipList.Random_level()
	// 如果是新增多层，每层头结点初始化
	if level > skipList.Level {
		for i := skipList.Level; i < level; i++ {
			update[i] = skipList.Header
		}
		skipList.Level = level
	}
	newNode := NewNode(key, level)
	for i := 0; i < level; i++ {
		// 前序节点的下一节点
		newNode.Forward[i] = update[i].Forward[i]
		// 更新前序节点的下一节点
		update[i].Forward[i] = *newNode
	}

}

func (skipList *SkipList) Random_level() int {
	level := 1
	// 返回一个int32类型的非负的31位伪随机数。
	// const SKIPLIST_P = 4
	for (rand.Int31()&0xFFFF)%SKIPLIST_P == 0 {
		level += 1
	}
	// const SKIPLIST_MAXLEVEL = 32
	if level < SKIPLIST_MAXLEVEL {
		return level
	} else {
		return SKIPLIST_MAXLEVEL
	}
}

func (skipList *SkipList) PrintSkipList() {
	fmt.Println("\nSkipList-------------------------------------------")
	for i := SKIPLIST_MAXLEVEL - 1; i >= 0; i-- {
		fmt.Println("level:", i)
		node := skipList.Header.Forward[i]
		for {
			if node.Value != nil {
				fmt.Printf("%d[%d:%p] ", node.Value.(int), i, &node.Value)
				node = node.Forward[i]
			} else {
				break
			}
		}
		fmt.Println("\n--------------------------------------------------------")
	} //end for

	fmt.Println("Current MaxLevel:", skipList.Level)
}

func (skipList *SkipList) Search(key int) *Node {

	node := skipList.Header
	for i := skipList.Level - 1; i >= 0; i-- {
		fmt.Println("\n Search() Level=", i) // 逐层检索
		for {
			if node.Forward[i].Value == nil {
				break
			}

			fmt.Printf("  %d ", node.Forward[i].Value)
			if node.Forward[i].Value.(int) == key {
				fmt.Println("\nFound level=", i, " key=", key)
				return &node.Forward[i]
			}

			if node.Forward[i].Value.(int) < key {
				node = &node.Forward[i]
				continue
			} else { // > key
				// 如果key要大，需要下沉，这里下沉复用上层的位置作为起始位置了么？
				// 如果没有，那就是辣鸡。。
				break
			}
		} //end for find

	} //end level
	return nil
}

func (skipList *SkipList) Remove(key int) {

	update := make(map[int]*Node)
	node := skipList.Header
	for i := skipList.Level - 1; i >= 0; i-- {
		for {

			if node.Forward[i].Value == nil {
				break
			}

			if node.Forward[i].Value.(int) == key {
				// 定位到每一层，该元素的位置，然后再逐一删除
				// 这里定位到的是，需要删除元素的，前置节点
				fmt.Println("Remove() level=", i, " key=", key)
				update[i] = node
				break
			}

			if node.Forward[i].Value.(int) < key {
				node = &node.Forward[i]
				continue
			} else { // > key
				break
			}

		} //end for find

	} //end level

	// 遍历待删除的每层节点（前置节点）
	for i, v := range update {
		// 如果节点就是头结点，则减少一层，同时删除
		if v == skipList.Header {
			skipList.Level--
		}
		// 带头结点的链表删除
		// 前置节点的下一跳指向待删除节点的下一跳
		v.Forward[i] = v.Forward[i].Forward[i]
	}
}
