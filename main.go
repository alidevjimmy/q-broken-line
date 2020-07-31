package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	final := make([]float64, n)
	for i := 0; i < n; i++ {
		res := .0
		scanner.Scan()
		n, _ := strconv.Atoi(strings.Split(scanner.Text(), " ")[0])
		for ii := 0; ii < n; ii++ {
			scanner.Scan()
			chars := scanner.Text()
			for _, char := range chars {
				if char == '\\' || char == '/' {
					res += .5 + float64(n-1-ii)
				} else if char == '_' {
					res += float64(n - 1 - ii)
				}
			}
		}
		final[i] = res
	}
	for _, val := range final {
		fmt.Printf("%.3f\n", val)
	}
}
