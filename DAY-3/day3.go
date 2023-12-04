package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	dt:=reader()
	//fmt.Println(dt)
	partOne(dt)
}

func reader() [][]rune {
	input := [][]rune{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, []rune(scanner.Text()))
	}
	return input
}

func partOne(data [][]rune){
	max:=0
	for i := range data {
		concat:=[]rune{}
		c:=false
		for j :=0; j<len(data[i]);j++ {
			el:=(data[i][j])
			if unicode.IsNumber(el) {
				concat=append(concat, el)

				if checkPose(data,i,j){
				c = true
				}

			}else{

				if len(concat)>0 && c{
					c=false
					for _, v := range concat {
						fmt.Print(string(v))
					}
					fmt.Println("")
					max+=int(convertRuneArrayToNumber(concat))
				}

				concat=[]rune{}
			}
			
		}
		
	}
	fmt.Println(max)
	
}

func convertRuneArrayToNumber(runes []rune) int64 {
	number, _ := strconv.ParseInt(string(runes), 10, 64)

	return number
  }

func checkPose(grid [][]rune, row, col int) bool {
	if row < 0 || row >= len(grid) {
		return false
	  }
	  
	  if col < 0 || col >= len(grid[0]) {
		return false
	  }
	// Controlla se il carattere in alto è diverso da un numero o da un punto
	if row > 0 && grid[row-1][col] != '.' && !unicode.IsNumber(grid[row-1][col]) {
	  return true
	}
  
	// Controlla se il carattere in basso è diverso da un numero o da un punto
	if row < len(grid)-1 && grid[row+1][col] != '.' && !unicode.IsNumber(grid[row+1][col]) {
	  return true
	}
  
	// Controlla se il carattere a sinistra è diverso da un numero o da un punto
	if col > 0 && grid[row][col-1] != '.' && !unicode.IsNumber(grid[row][col-1]) {
	  return true
	}
  
	// Controlla se il carattere a destra è diverso da un numero o da un punto
	if col < len(grid[0])-1 && grid[row][col+1] != '.' && !unicode.IsNumber(grid[row][col+1]) {
	  return true
	}
  
	// Controlla se il carattere in alto a sinistra è diverso da un numero o da un punto
	if row > 0 && col > 0 && grid[row-1][col-1] != '.' && !unicode.IsNumber(grid[row-1][col-1]) {
	  return true
	}
  
	// Controlla se il carattere in alto a destra è diverso da un numero o da un punto
	if row > 0 && col < len(grid[0])-1 && grid[row-1][col+1] != '.' && !unicode.IsNumber(grid[row-1][col+1]) {
	  return true
	}
  
	// Controlla se il carattere in basso a sinistra è diverso da un numero o da un punto
	if row < len(grid)-1 && col > 0 && grid[row+1][col-1] != '.' && !unicode.IsNumber(grid[row+1][col-1]) {
	  return true
	}
  
	// Controlla se il carattere in basso a destra è diverso da un numero o da un punto
	if row < len(grid)-1 && col < len(grid[0])-1 && grid[row+1][col+1] != '.' && !unicode.IsNumber(grid[row+1][col+1]) {
	  return true
	}
  
	// Nessun carattere adiacente è diverso da un numero o da un punto
	return false
  }
  