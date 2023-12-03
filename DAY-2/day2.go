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
	//partOne(fp)
	partTwo(fp)

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
		split:=strings.Split(line,":")
		split[1]=strings.Replace(split[1][1:],",","",-1)

		gameID,_:=strconv.Atoi(strings.Split(split[0]," ")[1])
		sim:=strings.Split(split[1],"; ")

		for _, v := range sim {
			simG:=strings.Split(v," ")

			for i := 0; i < len(simG); i+=2 {
				n,_:=strconv.Atoi(simG[i])
				if (mapp[simG[i+1]]<n){
					mapp[simG[i+1]]=n
				}
			
			}
			
		}
		fmt.Println(mapp,gameID)
		
		for tag := range mapp {
			if setup[tag]<mapp[tag]{	
				gameID=0
				break
			}
		}	 
		i+=gameID
	
	}
	fmt.Println(i)
}
func partTwo(input []string) {
	i:=0
	for _, line := range input {

		mapp:=make(map[string]int)
		split:=strings.Split(line,":")
		split[1]=strings.Replace(split[1][1:],",","",-1)

		gameID,_:=strconv.Atoi(strings.Split(split[0]," ")[1])
		sim:=strings.Split(split[1],"; ")

		for _, v := range sim {
			simG:=strings.Split(v," ")

			for i := 0; i < len(simG); i+=2 {
				n,_:=strconv.Atoi(simG[i])
				if (mapp[simG[i+1]]<n || mapp[simG[i+1]]==0){
					mapp[simG[i+1]]=n
				}
			
			}
			
		}
		fmt.Println(mapp,gameID)
		n:=1

		for tag := range mapp {

			n*=mapp[tag]
		}	 
		fmt.Println(n)
		i+=n
	
	}
	fmt.Println(i)
}