# Huffman Encoding Study

Welcome to the Huffman Encoding Study repository! This project benchmarks various data compression algorithms to evaluate their performance and efficiency. The key algorithms analyzed include:

- **Huffman Encoding**
- **Burrows-Wheeler Transform (BWT)**
- **Run-Length Encoding (RLE)**
- **Combined Approaches**: BWT + RLE + Huffman Encoding

## Features

- **Compression Ratio Calculation**: Measures how well each algorithm compresses data.
- **Performance Metrics**: Tracks compression and decompression times.
- **Space Efficiency**: Analyzes memory usage during compression and decompression.
- **Diverse Testing**: Evaluates algorithms with different input sizes and types.

## Getting Started

1. **Clone the Repository**:
```bash
git clone https://github.com/ml3m/huffman_encoding_study.git
```

Navigate to the Project Directory:

```bash
cd huffman_encoding_study
```
Build the Project:

```bash
go build -o compression_benchmark
```
Run the Benchmark:
```bash
./compression_benchmark [input_file]
```

Results
The benchmarking script compares:

Simple Encoding
Huffman Encoding
BWT + Huffman Encoding
RLE + Huffman Encoding
BWT + RLE + Huffman Encoding
Results include compression sizes, performance times, and compression ratios.

Contributing
Contributions are welcome! To contribute, please fork the repository and submit a pull request. For issues or feature requests, open an issue on GitHub.



### Key Sections:

1. **Overview**: Briefly describes what the project does.
2. **Features**: Lists the main features and metrics of the benchmarks.
3. **Getting Started**: Provides instructions on how to clone, build, and run the project.
4. **Results**: Describes what the benchmarking script evaluates.
6. **Contributing**: Information on how others can contribute to the project.

Huffman Coding in Go
Overview
Huffman coding is an algorithm used for lossless data compression. It assigns variable-length codes to input characters, with shorter codes assigned to more frequent characters. This program implements Huffman coding in Go, utilizing a priority queue to build the Huffman tree and visualizing the tree structure using the treedrawer library.

Table of Contents
Installation
Usage
Key Components
HuffmanNode
PriorityQueue
Functions
Huffman Tree Visualization
Example
License
Installation
Ensure you have Go installed on your machine. You can check by running:

bash
Copy code
go version
To install the treedrawer library, run:

bash
Copy code
go get github.com/m1gwings/treedrawer/tree
Usage
To run the program, use the following command:

bash
Copy code
go run main.go "<input string>"
Replace <input string> with the string you want to compress.

Key Components
HuffmanNode
The HuffmanNode struct represents a node in the Huffman tree. It contains:

char: The character represented by this node.
freq: The frequency of the character.
left: A pointer to the left child node.
right: A pointer to the right child node.
PriorityQueue
The PriorityQueue type is a slice of HuffmanNode pointers that implements a min-heap. It is used to efficiently build the Huffman tree based on character frequencies. It includes methods for:

Len(): Returns the length of the queue.
Less(): Defines the order of elements (by frequency and character).
Swap(): Swaps two elements in the queue.
Push(): Adds a new element to the queue.
Pop(): Removes and returns the highest priority element.
Functions
BuildHuffmanTree(input string) *HuffmanNode

Builds the Huffman tree from the input string and returns the root node.
GenerateHuffmanCodes(node *HuffmanNode, code string, codes map[rune]string)

Recursively generates Huffman codes for each character by traversing the tree.
DrawHuffmanTree(node *HuffmanNode, parent *tree.Tree, freqMap map[rune]int)

Visualizes the Huffman tree using the treedrawer library.
PrintFrequencyArray(freqMap map[rune]int)

Prints a sorted frequency array of characters in the input string.
Huffman Tree Visualization
The program visualizes the Huffman tree, displaying each character and its corresponding frequency at the leaf nodes. Internal nodes display the combined frequencies of their child nodes.

Example
Input
bash
Copy code
go run main.go "this is an example for huffman encoding"
Output
plaintext
Copy code
Frequency Array:
----------------
3 ( )
1 (a)
1 (d)
1 (e)
1 (f)
1 (g)
1 (h)
1 (i)
1 (l)
1 (m)
2 (n)
1 (o)
1 (r)
1 (s)
1 (t)
1 (u)
1 (x)

Huffman Codes:
--------------
Char  Code
      00
a     1111
d     1100
e     1110
f     1101
g     1010
h     0111
i     1000
l     1001
m     0101
n     0100
o     1011
r     0110
s     0010
t     0011
u     0010
x     1111

Huffman Tree
License
This project is licensed under the MIT License. See the LICENSE file for more information.

Feel free to modify any sections or add more information to fit your needs!
