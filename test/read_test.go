package test

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_all_read(t *testing.T) {
	fmt.Println("001")

	f, err := ioutil.ReadFile("../dictionary/SensitiveWords.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(f))
}

func Test_line_read(t *testing.T) {

	keywords := make([]string, 0)
	genre := make([]string, 0)

	fi, err := os.Open("../dictionary/SensitiveWords.txt")
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
		//fmt.Println(string(a))

		genre = append(genre, strings.Split(string(a), "|")[0])
		keywords = append(keywords, strings.Split(string(a), "|")[1])
	}

	for i := 0; i < len(keywords); i++ {
		fmt.Println(keywords[i])
	}

}
