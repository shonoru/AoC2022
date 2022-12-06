package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Parse(file io.Reader) []string {
	scanner := bufio.NewScanner(file)

	arr := make([]string, 0)

	for scanner.Scan() {
		var line = scanner.Text()
		arr = append(arr, line)
	}

	return arr
}

func CountCalories(arr []string) int {
	var total = 0
	var sum = 0

	for _, str := range arr {
		if str == "" {
			// nil

			if total < sum {
				fmt.Printf("cond %v < %v\n", total, sum)
				total = sum
			}

			sum = 0 // sum suppose to be re-setted
		} else {
			num, err := strconv.Atoi(str)

			if err == nil {
				sum += num
			}
			//fmt.Printf("%v\n", num)
		}
	}

	return total
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var arr = Parse(file)
	var ans = CountCalories(arr)
	fmt.Println(ans)
}
