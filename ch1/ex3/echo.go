package echo

import (
	"os"
	"strings"
)

// Echo3 concat all arguments using string#Join
func Echo3() {
	// fmt.Println(strings.Join(os.Args[1:], " "))
	strings.Join(os.Args[1:], " ")
}

// Echo2 concat all arguments using for ... range
func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
	}
	// fmt.Println(s)
}

// Echo1 concat ll arguments using for ...
func Echo1() {
	s, sep := "", ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	// fmt.Println(s)
}
