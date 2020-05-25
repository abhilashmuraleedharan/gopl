/*
 * As we saw, indexing a string yields its bytes, not its characters:
 * a string is just a bunch of bytes. That means that when we store a character value in a string,
 * we store its byte-at-a-time representation.
 * Let's look at a more controlled example to see how that happens.
 */
package main

import "fmt"

func main() {
	const placeOfInterest = `⌘`

	fmt.Printf("Type of placeOfInterest : ")
	fmt.Printf("%T\n", placeOfInterest)
	fmt.Printf("Plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n\n")

	// A for range loop, decodes one UTF-8-encoded rune on each iteration.
	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		// %#U verb displays both the code point's Unicode value and its printed representation
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}

/*
 * Source code in Go is defined to be UTF-8 text; no other representation is allowed.
 * When in source code we write the text `⌘`, the text editor used to create the program (VS code here)
 * places the UTF-8 encoding of the symbol ⌘ into the source text.
 * When we print out the hexadecimal bytes, we're just dumping the data the editor placed in the file.
 *
 * To summarize, strings can contain arbitrary bytes, but when constructed from string literals,
 * those bytes are (almost always) UTF-8.
 */
