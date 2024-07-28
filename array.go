/*
Sometimes things aren’t clear right away. That’s where you need to be patient and persevere and see where things lead.

– Mary Pierce
-
package main

import (

	"fmt"

)

func main() {

	//grades := [...]int{23, 45, 67, 87, 54, 31} // [...] it is called elapses where we do not give any value for arry length

	var grades [6]int = [6]int{23, 45, 67, 87, 54, 31}

	for index, elements := range grades {

		fmt.Println(index, "==>", elements)
	}

}
-

	package main

import "fmt"

	func main() {
	        arr := [...]string{1, 2, 3, 4, 5}
	        fmt.Println(arr)
	}

Ans: Error

-
package main

import "fmt"

	func main() {
		arr := [5]bool{true, true, true, true}
		for i := 0; i < len(arr); i++ {
			if arr[i] {
				fmt.Println(i)
			}
		}
	}

-

package main

import "fmt"

	func main() {
		arr := [5]bool{true, true}
		fmt.Println(arr)
	}

-
package main

import "fmt"

	func main() {
		arr := [4]int{10, 20, 30, 50, 90}
		fmt.Println(arr)
	}

-
package main

import "fmt"

	func main() {
		arr := [10]int{10, 20, 30, 50}
		fmt.Println(arr)
		fmt.Println(len(arr))
	}

-

package main

import "fmt"

	func main() {
		arr := [10]int{10, 20, 30, 50}
		fmt.Println(arr[0])
		fmt.Println(arr[2])
		fmt.Println(arr[4])
		fmt.Println(arr[8])
		fmt.Println(arr[10])
	}
*/
package main

import "fmt"

func main() {
	arr := [5][2]string{{"a"}, {"b"}, {"c"}}
	fmt.Println(arr[0][0])
	fmt.Println(arr[1][1])
	fmt.Println(arr[2][0])
}
