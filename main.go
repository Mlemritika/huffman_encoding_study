package main

import (
	"fmt"
	"time"
)

func main() {
	inputs := loadInputFromFile()

	fmt.Println("Performance and Space comparison for various inputs:")

    // BWT + RLE + Huffman Encoding

    //    BWT is:$: 1 bits                                                                                                                                
    //    RLE is:1$: 2 bits                                                                                                                               
    //    Huffman Encoding took 51.333µs                                                                                                                  
    //    BWT + RLE Compressed + huffman is:01: 2 bits

    start := time.Now()
    bwt := BWTTransform("")
    fmt.Printf("BWT is:%s: %d bits\n", bwt, MeasureSpace(bwt))
    compressedRLE := RLECompress(bwt)
    fmt.Printf("RLE is:%s: %d bits\n", compressedRLE, MeasureSpace(compressedRLE))
    root := BuildHuffmanTree(compressedRLE)
    codes := make(map[rune]string)
    GenerateCodes(root, "", codes)
    huffmanEncoded := Encode(compressedRLE, codes)
    TrackTime(start, "Huffman Encoding")
    huffmanSpace := MeasureSpace(huffmanEncoded)
    fmt.Printf("BWT + RLE Compressed + huffman is:%s: %d bits\n",huffmanEncoded, huffmanSpace)
    fmt.Printf("*****************************************************\n")
    fmt.Printf("*****************************************************\n")
    fmt.Printf("*****************************************************\n")
    fmt.Printf("*****************************************************\n")
    fmt.Printf("*****************************************************\n")
    fmt.Printf("*****************************************************\n")

	for _, in := range inputs {
        fmt.Printf("\n\n\n#####################################################")
        fmt.Println("\nInput:\n", in)
        fmt.Printf(      "#####################################################\n\n")


		// BWT + RLE + Huffman Encoding
        bwt := BWTTransform(in)
        fmt.Printf("BWT: %d bits\n", MeasureSpace(bwt))
		start := time.Now()
        compressedRLE := RLECompress(bwt)
        fmt.Printf("RLE: %d bits\n", MeasureSpace(compressedRLE))
		root := BuildHuffmanTree(compressedRLE)
		codes := make(map[rune]string)
		GenerateCodes(root, "", codes)
		huffmanEncoded := Encode(compressedRLE, codes)
		TrackTime(start, "Huffman Encoding")
		huffmanSpace := MeasureSpace(huffmanEncoded)
        fmt.Printf("BWT + RLE Compressed + huffman size: %d bits\n", huffmanSpace)
		fmt.Printf("*****************************************************\n")

		// RLE + Huffman Encoding
		start = time.Now()
        compressedRLE = RLECompress(in)
		root = BuildHuffmanTree(compressedRLE)
		codes = make(map[rune]string)
		GenerateCodes(root, "", codes)
		huffmanEncoded = Encode(compressedRLE, codes)
		TrackTime(start, "Huffman Encoding")
		huffmanSpace = MeasureSpace(huffmanEncoded)
        fmt.Printf("RLE Compressed + huffman size: %d bits\n", huffmanSpace)
		fmt.Printf("*****************************************************\n")

		// Classic Huffman Encoding
		start = time.Now()
		root = BuildHuffmanTree(in)
		codes = make(map[rune]string)
		GenerateCodes(root, "", codes)
		huffmanEncoded = Encode(in, codes)
		TrackTime(start, "Huffman Encoding")
		huffmanSpace = MeasureSpace(huffmanEncoded)
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

/* abandoned 

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
        */
	}
}
