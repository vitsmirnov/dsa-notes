package main

const bitCount int = 32

type BitTrieNode struct {
	children [2]*BitTrieNode
	count    int
}

func MakeBitTrieNode() *BitTrieNode {
	return &BitTrieNode{
		children: [2]*BitTrieNode{},
		count:    0}
}

type BitTrie struct {
	root *BitTrieNode
}

func MakeBitTrie() *BitTrie {
	return &BitTrie{root: MakeBitTrieNode()}
}

func (bt *BitTrie) Put(num int) {
	// time: O(b)
	// b - bit count

	node := bt.root
	node.count++
	for i := bitCount - 1; i >= 0; i-- {
		bit := (num >> i) & 1
		if node.children[bit] == nil {
			node.children[bit] = MakeBitTrieNode()
		}
		node = node.children[bit]
		node.count++
	}
}

func (bt *BitTrie) Remove(num int) {
	// time: O(b)
	// b - bit count

	node := bt.root
	node.count--
	for i := bitCount - 1; i >= 0; i-- {
		bit := (num >> i) & 1
		node = node.children[bit]
		node.count--
	}
}

func (bt *BitTrie) MaxXorWith(num int) int {
	// time: O(b)
	// b - bit count

	maxXor := 0
	node := bt.root
	for i := bitCount - 1; i >= 0; i-- {
		bit := ((num >> i) & 1) ^ 1
		if node.children[bit] != nil && node.children[bit].count > 0 {
			maxXor |= 1 << i
			node = node.children[bit]
		} else { // ~
			node = node.children[bit^1]
		}
	}
	return maxXor
}
