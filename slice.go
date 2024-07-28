/*
The secret of success is to do the common things uncommonly well.

– John D. Rockefeller
-
package main

import "fmt"

	func main() {
		//arr := []int{-1, -2}
		//for _, value := range arr {
		//	fmt.Println(value)
		//}
		arr := [5]string{"a", "b", "c", "d", "e"}
		slice := arr[:4]
		fmt.Println(slice)
		fmt.Println(len(slice))
		fmt.Println(cap(slice))

}
-
package main

import "fmt"

	func main() {
		arr := [5]string{"a", "b", "c", "d", "e"}
		slice := arr[:4]
		fmt.Println(arr)
		fmt.Println(slice)
		slice[1] = "x"
		fmt.Println(arr)
		fmt.Println(slice)

}

-
- create an integer array of size 5 with values [10, 20, 90, 70, 60]
- create a slice referencing the above array to contain elements from index 0, 1, 2
- change the element at index 2 of slice to 900

- print the array
- print the slice
-
package main

import "fmt"

	func main() {
		arr := [5]int{10, 20, 90, 70, 60}
		slice := arr[:3]
		fmt.Println(arr)
		fmt.Println(slice)
		slice[2] = 900
		fmt.Println(arr)
		fmt.Println(slice)
	}

-

package main

import "fmt"

	func main() {
		arr := [5]int{10, 20, 90, 70, 60}
		slice := arr[:3]
		fmt.Println(cap(slice))
		new_slice := append(slice, 100, 200)
		fmt.Println(cap(new_slice))
	}

-

package main

import "fmt"

	func main() {
		arr := [5]int{10, 20, 90, 70, 60}
		slice := arr[:3]
		fmt.Println(cap(slice))

		slice_2 := make([]int, 5, 20)
		new_slice := append(slice, slice_2...)
		fmt.Println(cap(new_slice))
	}

-
package main

import "fmt"

	func main() {
		arr := [5]int{10, 20, 90, 70, 60}
		slice := append(arr[:2], arr[3:])
		fmt.Println(slice)
	}

-

package main

import "fmt"

	func main() {
		arr := []int{10, 20, 90, 70, 60}
		slice := make([]int, 10)
		num := copy(slice, arr)
		fmt.Println(slice)
		fmt.Println(num)
	}

-

package main

import "fmt"

	func main() {
		arr := []int{10, 20, 90, 70, 60}
		slice := make([]int, 10)
		copy(slice, arr)
		slice[1] = 1000
		fmt.Println(arr)
		fmt.Println(slice)
	}
*/
package main

import "fmt"

func main() {
	arr := [10]int{10, 20}
	slice := arr[2:8]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}
