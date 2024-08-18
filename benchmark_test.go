package main

import (
	"testing"
)

// Benchmark for SimpleEncode
func BenchmarkSimpleEncode(b *testing.B) {
	in := "bbananabananabananabananabananabanbananaanaanana"
	for i := 0; i < b.N; i++ {
		simpleEncoded := SimpleEncode(in)
		MeasureSpace(simpleEncoded)
	}
}

// Benchmark for Huffman Encoding
func BenchmarkHuffmanEncode(b *testing.B) {
	in := "bbananabananabananabananabananabanbananaanaanana"
	for i := 0; i < b.N; i++ {
		huf := huffman_string(in)
		MeasureSpace(huf)
	}
}

// Benchmark for BWT + Huffman Encoding
func BenchmarkBWTHuffmanEncode(b *testing.B) {
	in := "bbananabananabananabananabananabanbananaanaanana"
	for i := 0; i < b.N; i++ {
		bwt := BWT_string(in)
		bwt_huf := huffman_string(bwt)
		MeasureSpace(bwt_huf)
	}
}

// Benchmark for RLE + Huffman Encoding
func BenchmarkRLEHuffmanEncode(b *testing.B) {
	in := "bbananabananabananabananabananabanbananaanaanana"
	for i := 0; i < b.N; i++ {
		rle := RLE_string(in)
		rle_huf := huffman_string(rle)
		MeasureSpace(rle_huf)
	}
}

// Benchmark for BWT + RLE + Huffman Encoding
func BenchmarkBWTRLEHuffmanEncode(b *testing.B) {
	in := "bbananabananabananabananabananabanbananaanaanana"
	for i := 0; i < b.N; i++ {
		bwt := BWT_string(in)
		bwt_rle := RLE_string(bwt)
		bwt_rle_huf := huffman_string(bwt_rle)
		MeasureSpace(bwt_rle_huf)
	}
}
