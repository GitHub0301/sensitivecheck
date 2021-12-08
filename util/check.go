package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Sensitivewords() (genre []string, keywords []string) {

	fi, err := os.Open("./dictionary/SensitiveWords.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		arr := strings.Split(string(a), "|")

		if len(arr) > 1 {
			genre = append(genre, arr[0])
			keywords = append(keywords, arr[1])
		}

	}

	return genre, keywords
}
