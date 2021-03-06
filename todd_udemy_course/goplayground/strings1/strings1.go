package main

import "fmt"

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("Println:")
	fmt.Println(sample)
	fmt.Printf("\n")

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n\n", sample)

	// %q (quoted) verb will escape any non-printable byte sequences
	// in a string so the output is unambiguous.
	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n\n", sample)

	/*
	 * If we are unfamiliar or confused by strange values in the string,
	 * we can use the "plus" flag to the %q verb. This flag causes the output
	 * to escape not only non-printable sequences, but also any non-ASCII bytes,
	 * all while interpreting UTF-8. The result is that it exposes the Unicode values
	 * of properly formatted UTF-8 that represents non-ASCII data in the string
	 */
	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)
}
