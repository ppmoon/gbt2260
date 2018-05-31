package gbt2260

//节点
type Node struct {
	value    string
	children map[string]*Node
}

//跟树
type Trie struct {
	root *Node
}

//创建一颗新树
func New() *Trie {
	return &Trie{
		root: &Node{
			children: make(map[string]*Node),
		},
	}
}

//返回跟节点
func (t *Trie) Root() *Node {
	return t.root
}

//添加节点
func (t *Trie) Add(lCode []string, name string) *Node {
	node := t.root
	for i := range lCode {
		r := lCode[i]
		if n, ok := node.children[r]; ok {
			node = n
		} else {
			//	否则就创建这个节点
			node = node.NewChild(r, name)
		}
	}
	return node
}

//创建并返回一个新子节点的指针这里的key
func (n *Node) NewChild(key string, value string) *Node {
	node := &Node{
		value:    value,
		children: make(map[string]*Node),
	}
	n.children[key] = node
	return node
}

//根据树遍历

// 返回一个子叶
func (n Node) Children() map[string]*Node {
	return n.children
}
