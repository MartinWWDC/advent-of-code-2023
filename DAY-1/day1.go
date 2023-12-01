package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
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

func sumArr(arr []int)int{
	n:=0
	for _, i := range arr {
		fmt.Println(i)
		n+=i
	}
	return n

}