// Time C. O(n), Space C. O(n)
// Runtime 719ms

type TrieNode struct {
	children map[rune]*TrieNode;
	isEnd bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

func (t *Trie) Insert(word string){
	node := t.root
	for _, c := range word {
		if node.children[c] == nil {
			node.children[c] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		node = node.children[c]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, c := range word {
		if node.children[c] == nil { return false }
		node = node.children[c]
	}
	return node.isEnd
}

func replaceWords(dictionary []string, sentence string)(res string) {
	minl := 99999999
	root := NewTrie()
	for _, word := range dictionary {
		minl = min(minl, len(word))
		root.Insert(word)
	}
	sWords := strings.Split(sentence, " ")
	for ii := 0; ii < len(sWords); ii++ {
		for i := minl; i < len(sWords[ii]); i++ {
			if root.Search(sWords[ii][:i]){
				sWords[ii] = sWords[ii][:i]
				break
			}
		}
	}
	res = strings.Join(sWords, " ")
	return
}