/*
The way to get started is to quit talking and begin doing.

– Walt Disney
-
package main

import "fmt"

func main() {
	for {
		fmt.Println("Hello World!") // it will result in infinite loop
	}
}

-
package main

import "fmt"

func main() {
        i := 3
        for i > 10 {
                fmt.Println(i * 2) //no out put as the condition in never met
                i += 1
        }
}

-
package main

import "fmt"

func main() {
        i := 5
        j := 0
        for j < 5 {
                fmt.Println(i * 2)
                j += 1
        }
}

-

package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i * i)
		if i == 3 {
			continue
		}
	}
}
*/

package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i * i)
		if i == 3 {
			break
		}
	}
}
