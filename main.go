package main

import (
	"bp-tree/src/tree"
	"fmt"
)

func main() {
	tree := tree.NewTree("abc.txt", 3, 512)

	// countries and their capitals
	tree.Put([]byte("India"), []byte("New Delhi"))
	tree.Put([]byte("Australia"), []byte("Sydney"))
	tree.Put([]byte("USA"), []byte("Washington DC"))
	tree.Put([]byte("Nepal"), []byte("Kathmandu"))
	tree.Put([]byte("Sri Lanka"), []byte("Colombo"))
	tree.Put([]byte("Bhutan"), []byte("Thimpu"))
	tree.Put([]byte("Pakistan"), []byte("Islamabad"))
	tree.Put([]byte("Zimbabwe"), []byte("Harare"))
	fmt.Println()

	// get India's capital
	fmt.Println("Attempting to fetch the capital of India")
	capital, err := tree.Get([]byte("India"))
	if err != nil {
		fmt.Printf("error while retrieving India's capital: %s\n", err)
	} else {
		fmt.Printf("India's capital: %s\n", string(capital))
	}
	fmt.Println("---------------------------------")
	fmt.Println()

	// get England's capital
	fmt.Println("Attempting to fetch the capital of England")
	capital, err = tree.Get([]byte("England"))
	if err != nil {
		fmt.Printf("error while retrieving England's capital: %s\n", err)
	} else {
		fmt.Printf("England's capital: %s\n", string(capital))
	}
	fmt.Println("---------------------------------")
	fmt.Println()

	// print all entries
	fmt.Println("Printing all entries in the tree")
	kvPairs := tree.Scan()
	for _, kv := range kvPairs {
		fmt.Printf("Country: %s, Capital: %s\n", string(kv.Key), string(kv.Value))
	}
	fmt.Println("---------------------------------")
	fmt.Println()

	// Correct Australia's capital
	fmt.Println("Correcting Australia's capital to Canberra")
	tree.Put([]byte("Australia"), []byte("Canberra"))
	fmt.Println("---------------------------------")
	fmt.Println()

	// Remove Pakistan's capital
	fmt.Println("Deleting entry corresponding to Pakistan")
	tree.Delete([]byte("Pakistan"))
	fmt.Println("---------------------------------")
	fmt.Println()

	// print all entries
	fmt.Println("Printing all entries in the tree")
	kvPairs = tree.Scan()
	for _, kv := range kvPairs {
		fmt.Printf("Country: %s, Capital: %s\n", string(kv.Key), string(kv.Value))
	}
	fmt.Println("---------------------------------")
	fmt.Println()

	// get capitals of all countries whose name lies between Barbados and New Zealand
	fmt.Println("Printing all entries lying between Barbados and Venezuela")
	kvPairs = tree.RangeQuery([]byte("Barbados"), []byte("Venezuela"))
	for _, kv := range kvPairs {
		fmt.Printf("Country: %s, Capital: %s\n", string(kv.Key), string(kv.Value))
	}
	fmt.Println("---------------------------------")
	fmt.Println()
}
