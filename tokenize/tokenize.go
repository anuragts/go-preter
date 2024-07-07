package tokenize

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(){
	file, err := os.Open("input.txt")
	if err != nil{
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil{
		fmt.Println(err)
	}
}

