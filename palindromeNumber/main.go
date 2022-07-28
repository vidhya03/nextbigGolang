package main

import "fmt"

func main() {
	fmt.Println( isPalindrome(11))//,true),
	fmt.Println( isPalindrome(121))//, true),)
	fmt.Println( isPalindrome(1221))//, true),))
	fmt.Println( isPalindrome(12455421))// , true),))
	fmt.Println( isPalindrome(15421))// , false),))
	fmt.Println( isPalindrome(0 ))//, true),))
	fmt.Println( isPalindrome(9 ))//, true),))
	fmt.Println( isPalindrome(1000021))// , false),))
	fmt.Println( isPalindrome(-29092))// , false),))
	fmt.Println( isPalindrome(1874994781))// , true)))
}

func isPalindrome(number int) bool {
	if(number <0){

		return  false;
	}
	var divider int = getDivider(number);
	for (divider > 1) {
		if ((number % 10) == (number / divider)) {
			number = number % divider;
			number = number / 10 ;
			divider = divider /100;
			continue;
		} else {
			return  false;
		}
	}
	return  true;
}
func getDivider(num int) int {
	var divider int = 1
	for (!((num/divider) < 10)) {
		divider *= 10
	}
	return  divider
}