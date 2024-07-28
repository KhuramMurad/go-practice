/*
Select the correct statement (s):

A. bitwise AND takes two numbers and does OR on every bit of two numbers.

B. bitwise OR takes two numbers and does OR on every bit of two numbers.

C. bitwise AND takes two numbers and does AND on every bit of two numbers.

D. bitwise OR takes two numbers and does AND on every bit of two numbers.

Answer = B,C
-

package main

import "fmt"

	func main() {
		var x, y int = 100, 90
		fmt.Println(x & y)
		fmt.Println(x | y)
	}

-

Select the correct statement(s):

A. The result of XOR is 0 when the two bits are same.

B. The result of XOR is 1 when the two bits are opposite.

C. The result of XOR is 0 when the two bits are opposite.

D. The result of XOR is 1 when the two bits are same.

Answer = A,bit

-

package main

import "fmt"

	func main() {
		var x, y int = 100, 90
		fmt.Println((x + y) >> 2) // in bitwise rightshift --> bits are shiftted to right in this case the bites woul be shifted to right
		// for 2 places and the remaining bits on left side are discarded
	}

-
*/
package main

import "fmt"

func main() {
	var x, y int = 100, 90
	fmt.Println(!(((x + y) >> 2) == 47))
}
