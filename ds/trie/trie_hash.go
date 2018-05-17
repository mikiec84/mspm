package trie

type TrieHashNode struct {
	Children map[byte]*TrieHashNode

	FinalState bool
}

func NewTrieHashNode() (tNode *TrieHashNode) {
	tNode = &TrieHashNode{Children: make(map[byte]*TrieHashNode), FinalState: false}
	return
}

func (root *TrieHashNode) Insert(key []byte) {
	var index byte

	tNode := root
	for level := 0; level < len(key); level++ {
		index = key[level]

		if tNode.Children[index] == nil {
			tNode.Children[index] = NewTrieHashNode()
		}
		tNode = tNode.Children[index]
	}

	tNode.FinalState = true
}

func (root *TrieHashNode) Search(key []byte) (found bool, final bool) {
	var index byte

	tNode := root
	for level := 0; level < len(key); level++ {
		index = key[level]

		if tNode.Children[index] == nil {
			return false, false
		}
		tNode = tNode.Children[index]
	}
	return tNode != nil, tNode.FinalState
}