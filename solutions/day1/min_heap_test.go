package day1

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	// Utworzenie kopca o rozmiarze 10
	h := NewHeap(10)

	// Test wstawiania elementów
	elements := []int{5, 3, 8, 1, 2, 9, 7}
	fmt.Println("Dodawanie elementów:", elements)
	for _, el := range elements {
		h.Insert(el)
		fmt.Println(h.String())
	}

	// Test wyciągania elementów (powinno wyciągać w kolejności rosnącej)
	fmt.Println("\nUsuwanie elementów:")
	for i := 0; i < len(elements); i++ {
		fmt.Printf("Wyciągnięty element: %d\n", h.Pop())
		fmt.Println(h.String())
	}

	// Test gdy kopiec jest pusty
	fmt.Println("\nPróba wyciągania z pustego kopca:")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Oczekiwany błąd:", r)
		}
	}()
	h.Pop() // Powinien spowodować błąd
}
