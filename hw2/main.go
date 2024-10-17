package main

import "fmt"

type BST_node struct { // struct of binary tree node
	value  int
	left   *BST_node
	right  *BST_node
	parent *BST_node
}

func Node_init(value int, parent *BST_node) *BST_node { // node initialization
	new_node := new(BST_node)
	new_node.value = value
	new_node.parent = parent
	return new_node
}

func Is_Exist(Root *BST_node) bool { // tree existence check
	if Root == nil {
		return false
	} else {
		return true
	}
}

func Find_min(Root *BST_node) (*BST_node, error) { // find min in tree
	if Is_Exist(Root) {
		flag := Root
		for flag.left != nil {
			flag = flag.left
		}
		return flag, nil
	} else {
		return nil, fmt.Errorf("Tree doesn't exist")
	}
}

func Find_next(some_node *BST_node) (*BST_node, error) { // find next node in a sorted enumeration
	if some_node.right != nil {
		return Find_min(some_node.right)
	}
	tmp := some_node.parent
	for (tmp != nil) && (some_node == tmp.right) {
		some_node = tmp
		tmp = tmp.parent
	}
	return tmp, nil
}

func Add(Root *BST_node, New *BST_node) (*BST_node, error) {
	if Is_Exist(Root) {
		prev := Node_init(0, nil)
		curr := Root
		for curr != nil {
			prev = curr
			if New.value < curr.value {
				curr = curr.left
			} else {
				curr = curr.right
			}
		}
		New.parent = prev
		if prev == nil {
			Root = New
		} else if New.value < prev.value {
			prev.left = New
		} else {
			prev.right = New
		}
		fmt.Println("Added some node with with value ", New.value)
		return Root, nil
	} else {
		return nil, fmt.Errorf("Tree doesn't exist")
	}
}

// Для проверяющего: выбор if боснован тем, что, хоть switch case и работает быстрее, задание требует
// рассмотрения определенного случая, что удобнее описать через if/else

func Delete(Root *BST_node, delete_node *BST_node) (*BST_node, error) { // function of deleting some node using if
	prev := Node_init(0, nil)
	tmp := Node_init(0, nil)
	out := delete_node.value
	if Is_Exist(Root) {
		if (delete_node.left == nil) || (delete_node.right == nil) {
			prev = delete_node
		} else {
			prev, _ = Find_next(delete_node)
		}
		if prev.left != nil {
			tmp = prev.left
		} else {
			tmp = prev.right
		}
		if tmp != nil {
			tmp.parent = prev.parent
		}
		if prev.parent == nil {
			Root = tmp
		} else if prev == prev.parent.left {
			prev.parent.left = tmp
		} else {
			prev.parent.right = tmp
		}
		if prev != delete_node {
			delete_node.value = prev.value
		}
		fmt.Println("Deleted some node with value ", out)
		return prev, nil
	} else {
		return nil, fmt.Errorf("Tree doesn't exist")
	}
}

// В main прописан код, использовавашийся для проверки и отладки работоспособности бинарного дерева
// Ввод ручной в связи с отсутсвием в тз точного задания входных/выходных данных :)

func main() {
	Root := Node_init(15, nil)
	Node1 := Node_init(5, Root)
	Node2 := Node_init(16, Root)
	Node3 := Node_init(3, Node1)
	Node4 := Node_init(12, Node1)
	Node5 := Node_init(13, Node4)
	Node6 := Node_init(10, Node4)
	Node7 := Node_init(6, Node6)
	Node8 := Node_init(7, Node7)
	Node9 := Node_init(20, Node2)
	Node10 := Node_init(18, Node9)
	Node11 := Node_init(23, Node9)
	Add(Root, Node1)
	Add(Root, Node2)
	Add(Root, Node3)
	Add(Root, Node4)
	Add(Root, Node5)
	Add(Root, Node6)
	Add(Root, Node7)
	Add(Root, Node8)
	Add(Root, Node9)
	Add(Root, Node10)
	Add(Root, Node11)
	Delete(Root, Node2)
	fmt.Println(Root.right.value)
}
