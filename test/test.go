package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("args:")
	for i, v := range os.Args {
		fmt.Printf("%d: %s\n", i, v)
	}
	var s string
	for {
		fmt.Print(": ")
		fmt.Scanf("%s", &s)
		if s == "exit" {
			fmt.Println("ok")
			break
		}
		fmt.Printf("cmd: %s\n", s)
	}
}
