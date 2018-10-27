package main

import (
	"calc/simplemath"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("stsarting...")
	args := os.Args
	fmt.Println(args[0])

	if args == nil || len(args) < 2 {
		Usage()
		return
	}

	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("Usage: calc add <integer1><integer2>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("Usage: calc add <integer1><integer2>")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result:", ret)
	}
}

var Usage = func() {
	fmt.Println("Usage: calc command [arguments]...")
	fmt.Println("\n The commands are:\n\tadd\tAddtion of two values.\n\tsqrt\tSquare" +
		" root of a non-negative value.")

}
