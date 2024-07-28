/*
package main

import "fmt"

func main() {
	var x, y string = "foo", "bar"
	x += y // similar as x == y , it is = operator in mathematics
	fmt.Println(x)
}
-
package main

import "fmt"

func main() {
        var x, y int = 27, 7
        x /= y // x = x / y --> x = 27/7
        fmt.Println(x)
}

-

package main

import "fmt"

func main() {
        var x, y float64 = 27.9, 7.0
        x -= y
        fmt.Println(x)
        x += y
        fmt.Println(x)
}

*/

package main

import "fmt"

func main() {
	var x, y int = 100, 9
	x /= y
	fmt.Println(x)
	x %= y
	fmt.Println(x)
}
