package main

// importing the packages
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile() {

	fmt.Printf("Writing to a file in Go lang\n")
	file, err := os.Create("files.txt")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	// Defer is used for purposes of cleanup
	defer file.Close()

	// len variable captures the length of the string written to the file.
	len, err := file.WriteString("Welcome to Kloudone." +
		" This program demonstrates reading and writing" +
		" operations to a file in Go lang.")

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	// Name() method returns the name of the
	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

func ReadFile() {

	fmt.Printf("\n\nReading a file in Go lang\n")
	fileName := "files.txt"

	data, err := ioutil.ReadFile("files.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)

}

// main function
func main() {

	CreateFile()
	ReadFile()
}