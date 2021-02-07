package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var age int
	var namn string
	var err error

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Vad heter du?")
	namn, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(namn)
	}

	fmt.Println("Hur gammal är du?")
	fmt.Scan(&age)

	if age > 48 {
		fmt.Println("Åh vad gammal du är")
	} else {
		fmt.Println("Ung")
	}

	var i int
	for i = 0; i < age; i++ {
		fmt.Printf("Varv %d\n", i)
	}
}
