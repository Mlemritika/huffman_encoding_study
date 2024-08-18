package main

import (
	"fmt"
	"time"
)

func main() {
	inputs := loadInputFromFile()

	fmt.Println("Performance and Space comparison for various inputs:")

	for _, in := range inputs {
        fmt.Printf("\n\n\n#####################################################")
        fmt.Println("\nInput:\n", in)
        fmt.Printf(      "#####################################################\n\n")

		// Classic Huffman Encoding
		start := time.Now()
		root := BuildHuffmanTree(in)
		codes := make(map[rune]string)
		GenerateCodes(root, "", codes)
		huffmanEncoded := Encode(in, codes)
		TrackTime(start, "Huffman Encoding")
		huffmanSpace := MeasureSpace(huffmanEncoded)
		fmt.Printf("Huffman Encoded Length: %d bits\n", huffmanSpace)
		fmt.Printf("*****************************************************\n")

		// Simple Encoding
		start = time.Now()
		simpleEncoded := SimpleEncode(in)
		TrackTime(start, "Simple Encoding")
		simpleSpace := MeasureSpace(simpleEncoded)
		fmt.Printf("Simple Encoded Length: %d bits\n", simpleSpace)
		fmt.Printf("Huffman Improvement Percentage: %.2f%%\n", getPercentage(huffmanSpace, simpleSpace))
        fmt.Printf("*****************************************************\n")


		// FGK Encoding
		fgkTree := NewFGKTree()
		start = time.Now()
		fgkEncoded := fgkTree.Encode(in)
		TrackTime(start, "FGK Encoding")
		fgkSpace := MeasureSpace(fgkEncoded)
		fmt.Printf("FGK Encoded Length: %d bits\n", fgkSpace)
		fmt.Printf("FGK Improvement Percentage: %.2f%%\n", getPercentage(fgkSpace, simpleSpace))
		fmt.Printf("\n#####################################################\n")
		fmt.Printf(  "#####################################################\n\n\n")
	}
}
