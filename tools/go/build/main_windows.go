// Build constraints, also known as build tags, control
// compilation by including or excluding files
// Compilation can also be controlled by the name of the
// file itself by "tagging" the file with a suffix
// (before the .go or .s extension)
// The file gopher_arm.go will only be compiled if the
// target processor is an ARM.

package main

func main() {
	println("helloworld")
}
