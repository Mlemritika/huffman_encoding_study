package main

import (
	"container/heap"
	"fmt"
	"os"
	"sort"
	"github.com/m1gwings/treedrawer/tree"
)

// Node structure for Huffman tree
type HuffmanNode struct {
	char  rune
	freq  int
	left  *HuffmanNode
	right *HuffmanNode
}

// Priority Queue implementation for a min-heap
type PriorityQueue []*HuffmanNode

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].freq == pq[j].freq {
		return pq[i].char < pq[j].char // Ensure consistent order by character
	}
	return pq[i].freq < pq[j].freq // Min-heap based on frequency
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*HuffmanNode))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

// BuildHuffmanTree builds the Huffman tree for the given input
func BuildHuffmanTree(input string) *HuffmanNode {
	freqMap := make(map[rune]int)
	for _, char := range input {
		freqMap[char]++
	}

	// Create a slice to hold the frequencies and characters
	type freqEntry struct {
		char rune
		freq int
	}
	frequencyEntries := make([]freqEntry, 0, len(freqMap))
	for char, freq := range freqMap {
		frequencyEntries = append(frequencyEntries, freqEntry{char: char, freq: freq})
	}

	// Sort the frequency entries first by frequency (descending), then by character (ascending)
	sort.Slice(frequencyEntries, func(i, j int) bool {
		if frequencyEntries[i].freq == frequencyEntries[j].freq {
			return frequencyEntries[i].char < frequencyEntries[j].char // Sort by ASCII value
		}
		return frequencyEntries[i].freq > frequencyEntries[j].freq // Sort by frequency descending
	})

	// Create a priority queue and add all characters to it
	pq := &PriorityQueue{}
	for _, entry := range frequencyEntries {
		heap.Push(pq, &HuffmanNode{char: entry.char, freq: entry.freq})
	}

	// Build the Huffman tree
	for pq.Len() > 1 {
		right := heap.Pop(pq).(*HuffmanNode) // Pop right (higher frequency)
		left := heap.Pop(pq).(*HuffmanNode)  // Pop left (lower frequency)
		newNode := &HuffmanNode{freq: left.freq + right.freq, left: left, right: right}
		heap.Push(pq, newNode)
	}

	return heap.Pop(pq).(*HuffmanNode) // Return the root of the tree
}

// GenerateHuffmanCodes generates Huffman codes for each character
func GenerateHuffmanCodes(node *HuffmanNode, code string, codes map[rune]string) {
	if node != nil {
		if node.left == nil && node.right == nil {
			codes[node.char] = code // Leaf node, store the code
		}
		GenerateHuffmanCodes(node.left, code+"0", codes)
		GenerateHuffmanCodes(node.right, code+"1", codes)
	}
}

// DrawHuffmanTree visualizes the Huffman tree using treedrawer
func DrawHuffmanTree(node *HuffmanNode, parent *tree.Tree, freqMap map[rune]int) {
	if node != nil {
		var current *tree.Tree
		if node.left == nil && node.right == nil {
			// Display character with frequency on leaf node
			current = parent.AddChild(tree.NodeString(fmt.Sprintf("%c (%d)", node.char, freqMap[node.char])))
		} else {
			current = parent.AddChild(tree.NodeInt64(int64(node.freq))) // Internal node
		}
		DrawHuffmanTree(node.left, current, freqMap)
		DrawHuffmanTree(node.right, current, freqMap)
	}
}

// PrintFrequencyArray prints the frequency array sorted by frequency and ASCII
func PrintFrequencyArray(freqMap map[rune]int) {
	fmt.Println("Frequency Array:")
	fmt.Println("----------------")
	
	type freqEntry struct {
		char rune
		freq int
	}
	
	frequencyEntries := make([]freqEntry, 0, len(freqMap))
	for char, freq := range freqMap {
		frequencyEntries = append(frequencyEntries, freqEntry{char: char, freq: freq})
	}

	sort.Slice(frequencyEntries, func(i, j int) bool {
		if frequencyEntries[i].freq == frequencyEntries[j].freq {
			return frequencyEntries[i].char < frequencyEntries[j].char // Sort by ASCII value
		}
		return frequencyEntries[i].freq > frequencyEntries[j].freq // Sort by frequency
	})

	for _, entry := range frequencyEntries {
		if entry.char == ' ' {
			fmt.Printf("%d (space)\n", entry.freq) // Handle space character separately
		} else {
			fmt.Printf("%d (%c)\n", entry.freq, entry.char)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go \"<input string>\"")
		return
	}

	input := os.Args[1] // Take input string as the first argument
	root := BuildHuffmanTree(input)

	// Frequency map
	freqMap := make(map[rune]int)
	for _, char := range input {
		freqMap[char]++
	}

	// Print frequency array
	PrintFrequencyArray(freqMap)

	// Generate Huffman codes
	codes := make(map[rune]string)
	GenerateHuffmanCodes(root, "", codes)

	// Print Huffman codes
	fmt.Println("Huffman Codes:")
	fmt.Println("--------------")
	fmt.Printf("%-5s %s\n", "Char", "Code")
	for char, code := range codes {
		fmt.Printf("%-5s %s\n", string(char), code)
	}
	fmt.Println()

	// Create a tree for drawing
	t := tree.NewTree(tree.NodeString("Huffman Tree"))
	DrawHuffmanTree(root, t, freqMap)

	// Draw the tree
	fmt.Println(t)
}
