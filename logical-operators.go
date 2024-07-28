/*
package main

import "fmt"

func main() {
        var a, b bool = true, false
        fmt.Println(a && b)
        fmt.Println(a || b)
}
-
package main

import "fmt"

func main() {

	var a, b bool = false, false
	fmt.Println(a && b)
	fmt.Println(a || b)
}
-
package main

import "fmt"

func main() {
        var a, b bool = false, true
        fmt.Println(!a)
        fmt.Println(b)
}

-

package main

import "fmt"

func main() {
        var a bool = false
        result := 10 > 50
        fmt.Println(!(a && result))
}

*/

package main

import "fmt"

func main() {
	var a bool = true
	result := 10 > 50
	fmt.Println(!(a || result))
}
