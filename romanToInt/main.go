package main

import (
	"fmt"
)

var romanBasic map[string]int


func init() {
    romanBasic = make(map[string]int)
	romanBasic["I"]= 1
	romanBasic["V"] = 5
	romanBasic["X"] = 10
	romanBasic["L"] = 50
	romanBasic["C"] = 100
	romanBasic["D"] = 500
	romanBasic["M"] = 1000  
}

  
func main() {

	romanLetters := [] string{"MCMLXXXIII","XXVI","III","LVIII", "MCMXCIV","X","IX"}

	for _, s := range romanLetters {
		fmt.Println(romanToInt (s))
	}
}

func romanToInt(s string) int {
	return romanToIntRecursive(s,0)
}

func romanToIntRecursive(romanLetters string, index int) int {

	 rLength := len(romanLetters);
	 total := 0
	if (rLength == 1 ) { return romanBasic[string(romanLetters[index])]}
	if(index < rLength && ((index +1 ) < rLength)){
		currentNumber := romanBasic[string(romanLetters[index])]
		nextNumber := romanBasic[string(romanLetters[index+1])] //check
		if(currentNumber >= nextNumber) {
			total += currentNumber;
			return total + romanToIntRecursive(romanLetters, index + 1);
		}else{
			total += (nextNumber - currentNumber);
			return total + romanToIntRecursive(romanLetters, index + 2);
		}
	} else {
		if ( 1 <= index && index < rLength) {
			return romanBasic[string(romanLetters[index])] // last number
		}else {
			return  0;
		}

	}
}
