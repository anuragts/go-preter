package tokenize

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


var keywords = map[string]bool{
	"if": true,
	"else": true,
	"for": true,
	"while": true,
	"do": true,
	"switch": true,
	"case": true,
	"default": true,
	"break": true,
	"continue": true,
	"return": true,
	"func": true,
	"package": true,
	"import": true,
}

// Over engineered if statement
func IsKeyword(word string) bool{
	return keywords[word]
}

func ReadFile(){
	file, err := os.Open("input.txt")
	if err != nil{
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	// Custom split function because I don't know if scanner class can split on space
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
        if atEOF && len(data) == 0 {
            return 0, nil, nil
        }
        if i := strings.IndexByte(string(data), ' '); i >= 0 {
            // We have a full token
            return i + 1, data[0:i], nil
        }
        // If we're at EOF, we have a final, non-empty, non-terminated token. Return it.
        if atEOF {
            return len(data), data, nil
        }
        // Request more data.
        return 0, nil, nil
    })

	for scanner.Scan(){
		fmt.Println(scanner.Text())
		fmt.Println(IsKeyword(scanner.Text()))
	}

	if err := scanner.Err(); err != nil{
		fmt.Println(err)
	}
}

