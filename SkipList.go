package SkipList

import (
"math/rand"
)

type SkipListNode struct {

    key int

    data int

    next []*SkipListNode

}
type SkipList struct {
	head *SkipListNode
	
	tail *SkipListNode
	
	length int

	level  int
    rand  *rand.Rand

}

func (list *SkipList) randomLevel() int { //随机生成的层数要满足P=0.5的几何分布

	level := 1
	
	for ; level < list.level && list.rand.Uint32() & 0x1 == 1; level ++{}
	
	return level
	
}

func (list *SkipList)insert(key int, data int)(bool){
	level := list.randomLevel()
	update := make([]*SkipListNode,level,level)
	node := list.head
	for index := level-1;index>=0;index--{
		for{
			node1 := node.next[index]
			if node1 == list.tail || node1.key > key{
				update[index] = node
				break
			} else if node1.key == key{
				node1.data = data
				return true
			} else{
					node = node1 // 往后查找
			}
		}
	}
	newNode := &SkipListNode{key, data, make([]*SkipListNode,level,level)}

	for index, node:=range update{
		node.next[index], newNode.next[index] = newNode,node.next[index]
	}
	
	list.length++
	return true
}

func (list *SkipList)delete(key int)bool{
	node := list.head
	remove := make([]*SkipListNode,list.level,list.level)
	var target *SkipListNode
	for index:=len(node.next)-1;index>=0;index--{
		for{
			node1 := node.next[index]
			if (node1 == nil || node1.key>key){
				break
			}else if (node1.key==key){
				remove[index] = node //需要更改next的元素（目标各层的前一个元素）
				target = node1
				break
			} else {
				node = node1
			}
		}
	}
	if target != nil{
		for index,node1:=range remove{ //此时node1为目标各层的前一个元素
			if node1 != nil{
				node1.next[index] = target.next[index] //更改各next
			}
		}
		list.length--
		return true
	}
	return false
}

func (list *SkipList)search(key int)int{
	node := list.head
	for index:=len(node.next)-1;index>=0;index--{
		node1 := node.next[index]
		for{
			if node1 == nil || node1.key>key{
				break
			}else if node1.key == key{
				return node1.data
			}else{
				node = node1
			}
		}
	}
	return -1
}
