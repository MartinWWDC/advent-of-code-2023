package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func main() {
	es2()
}

func es2(){
	arr:=part2()
	fmt.Println(arr)
}

func es1(){
	arr:=read()
	sum:=sumArr(arr)
	fmt.Println(sum)
}

func read() []int {
	scanner:=bufio.NewScanner(os.Stdin)
	arrInt:=[]int{}
	for scanner.Scan(){
		txt:=scanner.Text()
		arr:=[]string{}
		for _, v := range txt {
			if unicode.IsNumber(v){
				arr = append(arr, string(v))
			}
			
		}
		fmt.Println(arr[0],arr[len(arr)-1])
		nS:=arr[0]+arr[len(arr)-1]
		n,_:=strconv.Atoi(nS)
		arrInt=append(arrInt, n)
	}
	return arrInt
}
func part2() int {
	scanner:=bufio.NewScanner(os.Stdin)
	lines:=[]string{}
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}


	total := 0
	re := regexp.MustCompile("\\d|one|two|three|four|five|six|seven|eight|nine")

	for _, line := range lines {
		var current []string

		for i := range line {
			found := re.FindString(line[i:])
			if found == "" {
				continue
			}

			current = append(current, found)
		}

		if first, ok := digits[current[0]]; ok {
			total += first * 10
		}

		if last, ok := digits[current[len(current)-1]]; ok {
			total += last
		}
	}

	return total
}



func sumArr(arr []int)int{
	n:=0
	for _, i := range arr {
		fmt.Println(i)
		n+=i
	}
	return n

}