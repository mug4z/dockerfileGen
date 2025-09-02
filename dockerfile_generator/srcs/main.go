package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	test := flag.String("word", "TAMERE", "a string")
	test1 := flag.String("n", "TAMERE", "a string")
	test2 := flag.String("a", "TAMERE", "a string")
	is_write := flag.Bool("w", true, "if set write to file")
	argc := len(os.Args)
	if argc < 2 {
		os.Exit(1)
	}
	name := os.Args[1]
	flag.Parse()

	fmt.Println(name)
	fmt.Println("word: ", *test)
	fmt.Println("n: ", *test1)
	fmt.Println("a: ", *test2)
	if *is_write {
		d1 := []byte("PHILLIPPEEEE\n")
		os.WriteFile("GOTESTFILE", d1, 0644)
	}
}
