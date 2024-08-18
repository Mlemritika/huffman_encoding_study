package main

import (
	"fmt"
	"time"
	"strings"
	"math/rand"
	"errors"
)

import (
    "huffman_study/compression"
)


// Error handling for invalid data
func ValidateData(data string) error {
	if len(data) == 0 {
		return errors.New("input data is empty")
	}
	return nil
}

// Function to generate random test data
func GenerateRandomString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result strings.Builder
	for i := 0; i < length; i++ {
		result.WriteByte(chars[rand.Intn(len(chars))])
	}
	return result.String()
}

func huffmanString(input string) (string, int) {
	root := compression.BuildHuffmanTree(input)
	codes := make(map[rune]string)
	compression.GenerateCodes(root, "", codes)
	huffmanEncoded := compression.Encode(input, codes)
	return huffmanEncoded, MeasureSpace(huffmanEncoded)
}

func bwtString(input string) (string, int) {
	bwtEncoded := compression.BWTTransform(input)
	return bwtEncoded, MeasureSpace(bwtEncoded)
}

func rleString(input string) (string, int) {
	rleEncoded := compression.RLECompress(input)
	return rleEncoded, MeasureSpace(rleEncoded)
}

func main() {
	// Sample input and various test cases
    inputs := loadInputFromFile()

	fmt.Println("Performance and Space comparison for various inputs:")

	for _, in := range inputs {
		fmt.Printf("\nTesting with input size: %d characters\n", len(in))
        fmt.Printf("String:%s\n\n", in)

		// Simple Encoding
		start := time.Now()
		simpleEncoded := compression.SimpleEncode(in)
		simpleSize := MeasureSpace(simpleEncoded)
		TrackTime(start, "Simple Encoding")
		fmt.Printf("Simple Encoded size: %d bits\n", simpleSize)
		fmt.Printf("Compression Ratio: %.5f\n\n", CompressionRatio(len(in)*8, simpleSize))

		// Huffman Encoding
		start = time.Now()
		_, hufSize := huffmanString(in)
		TrackTime(start, "Huffman Encoding")
		fmt.Printf("Huffman Encoded size: %d bits\n", hufSize)
		fmt.Printf("Compression Ratio: %.5f\n\n", CompressionRatio(len(in)*8, hufSize))

		// BWT + Huffman Encoding
		start = time.Now()
		_, bwtHufSize := huffmanString(compression.BWTTransform(in))
		TrackTime(start, "BWT + Huffman Encoding")
		fmt.Printf("BWT + Huffman Encoded size: %d bits\n", bwtHufSize)
		fmt.Printf("Compression Ratio: %.5f\n\n", CompressionRatio(len(in)*8, bwtHufSize))

		// RLE + Huffman Encoding
		start = time.Now()
		_, rleHufSize := huffmanString(compression.RLECompress(in))
		TrackTime(start, "RLE + Huffman Encoding")
		fmt.Printf("RLE + Huffman Encoded size: %d bits\n", rleHufSize)
		fmt.Printf("Compression Ratio: %.5f\n\n", CompressionRatio(len(in)*8, rleHufSize))

		// BWT + RLE + Huffman Encoding
		start = time.Now()
		bwtRle := compression.RLECompress(compression.BWTTransform(in))
		_, bwtRleHufSize := huffmanString(bwtRle)
		TrackTime(start, "BWT + RLE + Huffman Encoding")
		fmt.Printf("BWT + RLE + Huffman Encoded size: %d bits\n", bwtRleHufSize)
		fmt.Printf("Compression Ratio: %.5f\n\n", CompressionRatio(len(in)*8, bwtRleHufSize))
        fmt.Println("#####################################################")
        fmt.Println("#####################################################")
	}

	fmt.Println("*****************************************************")
}
