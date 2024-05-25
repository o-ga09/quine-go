package main

import "fmt"

func main() {
	var s = "package main\n\n import \"fmt\"\n\n func main() {\n var s = %q \n fmt.Printf(s,s)\n }\n"
	fmt.Printf(s, s)
}
