package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	Replace = strings.Replace

	a string
	b string
)

func init() {
	flag.StringVar(&b, "b", "", "")
	flag.StringVar(&a, "a", "", "")
	flag.Parse()
}

func main() {
	if !flag.Parsed() {
		log.Fatal("flags not parsed")
	}

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		fmt.Println(Replace(scan.Text(), a, b, 1))
	}
	if err := scan.Err(); err != nil {
		log.Fatalf("scan error: %s", err)
	}
}
