package main

import (
	"fmt"
	"strings"
	"unicode"
)

// ReverseAlphabet mengembalikan string dengan urutan alphabet terbalik
func ReverseAlphabet(input string) string {

	var letters []rune
	var digits []rune

	// Pisahkan huruf dan angka
	for _, r := range input {
		if unicode.IsLetter(r) {
			letters = append(letters, r)
		} else if unicode.IsDigit(r) {
			digits = append(digits, r)
		}
	}

	// Balik urutan huruf
	for i, j := 0, len(letters)-1; i < j; i, j = i+1, j-1 {
		letters[i], letters[j] = letters[j], letters[i]
	}

	// Gabungkan huruf yang sudah dibalik dengan angka
	result := string(letters) + string(digits)
	return result
}

// LongestWord mengembalikan kata terpanjang dari sebuah kalimat
func LongestWord(sentence string) string {
	words := strings.Fields(sentence)
	longest := ""
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}
	return longest
}

// CountOccurrences menghitung berapa kali kata-kata dalam query muncul di input
func CountOccurrences(input, query []string) []int {
	counts := make([]int, len(query))
	inputMap := make(map[string]int)

	for _, word := range input {
		inputMap[word]++
	}

	for i, word := range query {
		counts[i] = inputMap[word]
	}

	return counts
}

// DifferenceOfDiagonals menghitung perbedaan antara jumlah diagonal matriks
func DifferenceOfDiagonals(matrix [][]int) int {
	var diagonal1, diagonal2 int

	for i := 0; i < len(matrix); i++ {
		diagonal1 += matrix[i][i]
		diagonal2 += matrix[i][len(matrix)-1-i]
	}

	return diagonal1 - diagonal2
}

func main() {
	// Reverse Alphabet
	reversed := ReverseAlphabet("NEGIE1")
	fmt.Println("Reversed Alphabet:", reversed) // Output: EIGEN1

	// Longest Word
	longest := LongestWord("Saya sangat senang mengerjakan soal algoritma")
	fmt.Println("Longest Word:", longest) // Output: mengerjakan

	// Count Occurrences
	input := []string{"xc", "dz", "bbb", "dz"}
	query := []string{"bbb", "ac", "dz"}
	counts := CountOccurrences(input, query)
	fmt.Println("Occurrences:", counts) // Output: [1 0 2]

	// Difference of Diagonals
	matrix := [][]int{{1, 2, 0}, {4, 5, 6}, {7, 8, 9}}
	diff := DifferenceOfDiagonals(matrix)
	fmt.Println("Difference of Diagonals:", diff) // Output: 3
}
