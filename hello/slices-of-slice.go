package main

import (
	"fmt"
	"strings"
	// "strings"
)

func main (){
   
	board := [][]string {
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}
	fmt.Println(board)

	board[0][0] = "x"
	board[1][1] = "x"
	board[2][2] = "x"


	for i:=0 ; i < len(board); i++ {
		fmt.Printf("%s\n",  board[i])
	}

	for i:=0 ; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i]," "))
	}

}