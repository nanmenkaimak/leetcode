package main

type Trie struct {
	children [26]*Trie
	isLast   bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	curr := this
	for i := range word {
		idx := word[i] - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &Trie{}
		}
		curr = curr.children[idx]
	}
	curr.isLast = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for i := range word {
		idx := word[i] - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return curr.isLast
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for i := range prefix {
		idx := prefix[i] - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
