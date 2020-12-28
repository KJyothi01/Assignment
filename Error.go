package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("files.txt") //Files.txt file is there in my system
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Failed to open file at path", pErr.Path)
			return
		}
		fmt.Println("Generic error", err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}