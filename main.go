package main

import "fmt"

func isPalindrome(value string) bool {
	//Gerin Aryo Prasetia - Aqua Developer Batch 1
	//Palindrom= character yang dibaca tetap sama baik dari depan atau belakang
	//Melakukan iterasi sebanyak length dari value yang diterima
	//Karena palindrome dibaca sama dari depan atau belakang
	//maka akan nilai di indeks ke-0 akan sama dengan nilai indeks ke-n-length-1 (var j di algo dibawah ini)
	//KOMPLEKSITAS ALGORITMA ; 0(n)
	for i := 0; i < len(value); i++ {
		j := len(value) - 1 - i
		//Iterasi akan berhenti jika terdapat character yang dibandingkan berbeda nilainya
		if value[i] != value[j] {
			return false
		}
	}
	return true
}

func main() {
	var tester = []string{"Gerin", "kasur ini rusak", "ini", "414", "xyzz"}
	for i := 0; i < len(tester); i++ {
		fmt.Printf("%v\t %v\n", tester[i], isPalindrome(tester[i]))
	}
}
