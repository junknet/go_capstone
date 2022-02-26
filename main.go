package main

import "fmt"

func main() {
	var code = "\x02\x90\x20\x46\xFF\xF7\xE2\xFD"
	input := []byte(code)
	Instructions := decompiler(input, 0)
	for _, i := range Instructions {
		fmt.Printf("%s %s\n", i.Mnemonic, i.OpStr)
	}
}
