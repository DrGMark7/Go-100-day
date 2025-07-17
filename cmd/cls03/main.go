package main

func main() {
	name := []string {"Alice", "Bob", "Charlie"}
	for _, n := range name {
		if n == "Bob" {
			continue
		}
		println("Hello, " + n)
	}
}
