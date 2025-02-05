/*
package main

import "fmt"

	func main() {
		var a, b string = "kolkata", "Kolkata"
		if a == b {
			fmt.Println("strings are equal")
		} else {
			fmt.Println("strings are not equal")
		}
		fmt.Println("thank you!")

}
-

package main

import "fmt"

	func main() {
	        var a, b string = "kolkata", "kolkata"
	        if a == b {
	                fmt.Println("strings are equal")
	        }
	        else {
	                fmt.Println("strings are not equal")
	        }
	        fmt.Println("thank you!")

}
-
package main

import "fmt"

	func main() {
	        var a, b string = "foo", "bar"
	        if a+b == "foo" {
	                fmt.Println("foo")
	        } else if a+b == "bar" {
	                fmt.Println("bar")
	        } else if a+b == "foobar" {
	                fmt.Println("foobar")
	        } else {
	                fmt.Println("None matched")
	        }
	        fmt.Println("thank you!")

}

// -
*/
package main

import "fmt"

func main() {
	var a, b string = "foo", "bar"
	if a+b == "foo" {
		fmt.Println("foo")
		//} else if a+b == "foobar" { // the conditionis true in this part of code, so this block of code is executed and exites the loop to lat line of code
		//	fmt.Println("bar")
	} else if a+b == "foobar" {
		fmt.Println("foobar")
	} else {
		fmt.Println("None matched")
	}
	fmt.Println("thank you!")

}

/*
Can we use if block independently, without else-if and else blocks ?


A. No, we need atleast one else block.


B. else-if is optional, but if block would raise error without else block.


C. Yes, we can write if block without any of the other blocks.


D. We need atleast one if, one else-if and one else block for program to compile successfully.

Solution = Can
*/
