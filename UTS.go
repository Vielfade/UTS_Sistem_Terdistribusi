package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		kata := strings.Split(scanner.Text(), " ")
		counting := make(map[string]int) 
    	for _, jumlahkata := range kata { 
        counting[jumlahkata] += 1 
    } 
    fmt.Println(counting)
	}
}