package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

const one = 1


func checkNum(num int) int {
	sum:=0
	for i:=1;i<=num/2;i++{
		if num%i==0{
			sum= sum+i
		}
	}
	if sum==num{
		return num
	}
	return 0
}





func find(numbers []int) []string {
	var s []string

	for _, num := range numbers {
		if num < one {
			s = append(s, fmt.Sprintf("%v - number must be positive", num))
			continue
		}
		for i := 1; i < num; i++ {
				numtoa:=checkNum(i)
				if numtoa!=0 {
					s = append(s, fmt.Sprintf("perfect num= %v", numtoa))
				}
		}

	}

	return s
}

func check(numbers []string) []int {
	nums:= make([]int, 0)

	for _, number := range numbers {
		nint, err := strconv.Atoi(number)
		if err != nil {
			continue
		}

		nums = append(nums, nint)
	}

	return nums
}

func input() []string {
	var n []string

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Input number: ")
		scanner.Scan()
		text := scanner.Text()
		n = append(n, text)

		if text == "" {
			n = n[0 : len(n)-1]

			break
		}
	}

	return n
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	var numbers []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}

	return numbers
}

func main() {
	filename := flag.String("f", "", "for taking numbers from file")

	flag.Parse()

	args := os.Args[1:]

	switch {
	case len(*filename) != 0:
		fmt.Println(find(check(readFile(*filename))))
	case len(args) == 0:
		fmt.Println(find(check(input())))
	default:
		fmt.Println(find(check(args)))
	}
}