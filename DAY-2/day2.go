package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var setup = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	//fmt.Println(setup)
	fp := reader()
	partOne(fp)

}

func reader()[]string{
	input:=[]string{}	
	scanner:=bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func partOne(input []string) {
	i:=0
	for _, line := range input {
		mapp:=make(map[string]int)
		split:=strings.Split(line," ")
		gameID,_:=strconv.Atoi(split[1][:len(split[1])-1])
		//fmt.Println(gameID)
		
		for i := 2; i < len(split); i+=2 {
			count,_:=strconv.Atoi(split[i])
			tag:=strings.Replace(split[i+1][:len(split[i+1])],",","",-1)
			
			mapp[tag]+=count
		}
		
		
		fmt.Println(mapp)
		for tag := range mapp {
			if setup[tag]<=mapp[tag]{
				
				//fmt.Println(tag,setup[tag],mapp[tag])
				gameID=0
				break
			}
		}
		fmt.Println(gameID)
		i+= gameID
	}
	fmt.Println(i)
}