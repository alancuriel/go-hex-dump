package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No filename passed as argument")
		os.Exit(1)
	}

	var fileName string = os.Args[1]

	file, err := os.Open(fileName)
	handleError(err)

	hexDumpFile(file)
}

func hexDumpFile(file *os.File) {
	var i int = 0
	buffer := make([]byte, 16)
	bytesRead, err := file.Read(buffer)

	handleError(err)

	for bytesRead != 0 {
		fmt.Printf("%07X | ", i)
		fmt.Printf("% X\n", buffer)

		clear(buffer)

		bytesRead, err = file.Read(buffer)

		i += 16
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func hexPrint(str string) {
	str_bytes := []byte(str)
	n := len(str_bytes)

	for i, b := range str_bytes {
		fmt.Printf("%X ", b)
		if (i+1)%16 == 0 || i+1 >= n {
			fmt.Println()
		}
	}
}
